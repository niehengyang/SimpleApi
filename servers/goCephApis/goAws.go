package goCephApis

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
)

type CephMgmt struct {
	host       string `ceph host`
	DoName     string `ceph doname`
	bucket_id  string `bucket_id`
	PathStyle  bool   `ceph url style, true means you can use host directly,  false means bucket_id.doname will be used`
	AccessKey  string `aws s3 aceessKey`
	SecretKey  string `aws s3 secretKey`
	block_size int64  `block_size`
}

var Ceph *CephMgmt

type MyProvider struct{}

func (m *MyProvider) Retrieve() (credentials.Value, error) {

	return credentials.Value{
		AccessKeyID:     Ceph.AccessKey,
		SecretAccessKey: Ceph.SecretKey,
	}, nil
}

func (m *MyProvider) IsExpired() bool { return false }

func (this *CephMgmt) Init() error {
	this.host = beego.AppConfig.String("ceph::host")
	if len(this.host) <= 0 {
		return errors.New("ceph conf host is nil")
	}
	this.bucket_id = beego.AppConfig.String("ceph::bucket_id")
	if len(this.bucket_id) <= 0 {
		return errors.New("ceph conf bucket_id is nil")
	}
	this.block_size, _ = beego.AppConfig.Int64("ceph::block_size")
	if this.block_size <= 5*1024*1024 {
		this.block_size = 5 * 1024 * 1024
	}
	return nil
}

func (this *CephMgmt) Init2(host string, bucket string, accesskey string, secretkey string) error {

	this.host = host
	this.bucket_id = bucket
	this.AccessKey = accesskey
	this.SecretKey = secretkey
	this.PathStyle = true
	this.block_size = 8 * 1024 * 1024

	return nil
}

func (this *CephMgmt) connect() (*s3.S3, error) {
	/** 创建连接**/
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("default"), //Required 目前尚未分区，填写default即可
			//EndpointResolver: endpoints.ResolverFunc(s3CustResolverFn),
			Endpoint:         &this.host,
			S3ForcePathStyle: &this.PathStyle,
			Credentials:      credentials.NewCredentials(&MyProvider{}),
		},
	}))
	// Create the S3 service client with the shared session. This will
	// automatically use the S3 custom endpoint configured in the custom
	// endpoint resolver wrapping the default endpoint resolver.
	return s3.New(sess), nil
}

/** 下载文件**/
func (this *CephMgmt) Download(src_name, dst_name string) error {
	s3Svc, _ := this.connect()
	// Operation calls will be made to the custom endpoint.
	resp, err := s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: &this.bucket_id, //Required 可认为是cpid，也可认为是对于根目录而言的第一级目录.必须先创建，在存储。
		Key:    &src_name,       //Required 文件名，中间可以带着路径，格式如：{path}/{filename}
		//Range:  aws.String("bytes=0-499"),             //not must be Required 文件范围，如果没有则是全文件
	})
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println("hello , i met error")
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("AcceptRanges :", *resp.AcceptRanges)
	fmt.Println("ContentLength:", *resp.ContentLength)
	//fmt.Println("ContentRange :", *resp.ContentRange)
	//fmt.Println("PartsCount   :", *resp.PartsCount)
	// Pretty-print the response data.
	//fmt.Println(resp)

	//数据都在resp.Body中，可以转储到文件，具体操作参考go的操作
	out, err := os.OpenFile(dst_name, os.O_CREATE|os.O_RDWR, 0666)
	if out == nil {
		fmt.Println("Open fail")
		return err
	}
	num, err := io.Copy(out, resp.Body)
	fmt.Printf("\n write %d err %v \n", num, err)
	return nil
}

/** 上传文件**/
func (this *CephMgmt) Upload(src_name, dst_name string) error {

	fmt.Println(this.bucket_id)

	s3Svc, _ := this.connect()

	dst_full_name := fmt.Sprintf("/%s/%s", this.bucket_id, dst_name) //Notice: 如果要讲桶ID拼到全路径下，一定要在最前面加'/'

	//////////////////////////////////////// 存入 /////////////////////////////
	/////////////// 创建一个分片上传context
	param_init := &s3.CreateMultipartUploadInput{
		Bucket: aws.String(this.bucket_id), // Required {bucket}
		Key:    aws.String(dst_full_name),  // Required /{bucket}/{path}/{filename}
		/*
		   Metadata: map[string]*string{
		       "Key": aws.String("lasttime_if_need"), // Required 可以填充一些文件属性,ceph云存储不关心其中的内容
		   },
		*/
	}

	resp_init, err := s3Svc.CreateMultipartUpload(param_init)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println("I am create upload.", err.Error())
		return err

	}
	// Pretty-print the response data.
	up_id := resp_init.UploadId
	fmt.Println(resp_init) // 需要保存该resp , resp.UploadId 是对象网关创建的标识

	///////////////////////// 分片上传
	f, err := os.Open(src_name)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	buf := make([]byte, this.block_size) //TODO:一次读取多少个字节

	var part_num int64 = 0
	var completes []*s3.CompletedPart

	for {
		n, err := bfRd.Read(buf)
		if err == io.EOF {
			fmt.Println("read data finished")
			break
		}
		if int64(n) != this.block_size {
			data := make([]byte, n)
			data = buf[0:n]
			buf = data
		}
		//////////////////////// 每次上传一个分片，每次的PartNumber都要唯一
		part_num++
		param := &s3.UploadPartInput{
			Bucket:        aws.String(this.bucket_id), // Required bucket
			Key:           aws.String(dst_name),       // Required {path}/{filename}
			PartNumber:    aws.Int64(part_num),        // Required 每次的序号唯一且递增
			UploadId:      up_id,                      // Required 创建context时返回的值
			Body:          bytes.NewReader(buf),       // Required 数据内容
			ContentLength: aws.Int64(int64(n)),        // Required 数据长度
		}

		resp2, err := s3Svc.UploadPart(param)
		if err != nil {
			fmt.Printf("Hello ,i am wrong[%s][%d][%d]\n", dst_name, part_num, n)
			return err
		}
		fmt.Println(resp2) // 需要保存该resp，因为resp.Etag 需要在通知完成上传时使用

		var c s3.CompletedPart
		c.PartNumber = aws.Int64(part_num) // Required Etag对应的PartNumber, 上一步返回的
		c.ETag = resp2.ETag                // Required 上传分片时返回的值 Etag
		completes = append(completes, &c)
	}

	/////////////////////////// 结束上传
	params := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(this.bucket_id), // Required {bucket}
		Key:      aws.String(dst_name),       // Required {path}/{filename}
		UploadId: up_id,                      // Required 创建context时返回的值
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completes,
		},
		//RequestPayer: aws.String("RequestPayer"),
	}
	resp_comp, err := s3Svc.CompleteMultipartUpload(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return err
	}
	// Pretty-print the response data.
	fmt.Println(resp_comp)
	return nil
}

/** 遍历目录**/
func (this *CephMgmt) DirectoryTraversal(src_name, dst_name string) (objkeys []string, err error) {

	return objkeys, nil
}
