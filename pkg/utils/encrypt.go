package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const md5key = "YourSecretEncryptKey"

/*
这个算法第一个字符是随机的从密钥中获取一个字符串，
最后一个字符是验证字符，是利用(密钥长度-随机数+给定ID长度)%密钥长度
中间部分就是：每位转换成16进制的ID数字+随机数，然后取余key的长度，得到的数字就是该位在key的字符，加入到总加密后的字符串中
*/
func Encode(id int64) (string, error) {
	// 16进制数组
	hexArr := strings.Split(strconv.FormatInt(id, 16), "")
	keyArr := strings.Split(md5key, "")

	keyLen := len(md5key)
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(keyLen - 1)
	str := keyArr[rnd]

	verify := keyArr[(keyLen-rnd+len(strconv.FormatInt(id, 10)))%keyLen]

	var buf bytes.Buffer
	buf.WriteString(str)
	for _, h := range hexArr {

		hi, err := strconv.ParseInt(h, 16, 64)
		if err != nil {
			return "", err
		}
		offset := int(hi) + rnd

		buf.WriteString(keyArr[offset%keyLen])
	}
	buf.WriteString(verify) // 写入验证码
	return buf.String(), nil
}

func Decode(str string) (int64, error) {

	strArr := strings.Split(str, "")
	keyArr := strings.Split(md5key, "")

	keyLen := len(md5key)

	rnd := strings.Index(md5key, strArr[0])
	verify := strArr[len(strArr)-1]

	strArr = strArr[1 : len(strArr)-1]

	var buf bytes.Buffer
	for _, s := range strArr {
		pos := strings.Index(md5key, s)
		if pos >= rnd {
			buf.WriteString(strconv.FormatInt(int64(pos-rnd), 16))
		} else {
			buf.WriteString(strconv.FormatInt(int64(keyLen-rnd+pos), 16))
		}
	}

	dec, err := strconv.ParseInt(buf.String(), 16, 64)
	if err != nil {
		return 0, err
	}

	if verify != keyArr[(keyLen-rnd+len(strconv.FormatInt(dec, 10)))%keyLen] {
		return 0, errors.New("校验错误，给定字符串不合法")
	}

	return dec, nil

}

/**_________________________高级加密标准（Adevanced Encryption Standard ,AES）____________________________**/

const aesKey = "123456781234567812345678"

func AesEncrypt(orig string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(aesKey)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error()))
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	return base64.RawURLEncoding.EncodeToString(cryted)

}

func AesDecrypt(cryted string) string {
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	crytedByte, _ := base64.RawURLEncoding.DecodeString(cryted)
	k := []byte(aesKey)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error()))
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/**_________________________美国数据加密标准(DES)____________________________**/

var desKey = []byte("2fa6c1e9")

//加密
func Encrypt(text string) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(desKey)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

//解密
func Decrypt(decrypted string) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(desKey)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
