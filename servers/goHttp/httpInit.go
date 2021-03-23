package goHttp

import (
	"github.com/kirinlabs/HttpRequest"
	"net"
	"net/http"
	"time"
)

var httpReq *HttpRequest.Request
var transport *http.Transport

func init() {
	httpReq = HttpRequest.NewRequest()
	transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

//头部
func SetHeader(req *HttpRequest.Request, headerData map[string]string) *HttpRequest.Request {
	req.SetHeaders(headerData)
	return req
}

//校验码
func SetCookies(req *HttpRequest.Request, cookieData map[string]string) *HttpRequest.Request {
	req.SetCookies(cookieData)
	return req
}

//身份
func SetBaseAuth(req *HttpRequest.Request, username string, password string) *HttpRequest.Request {
	req.SetBasicAuth(username, password)
	return req
}

//超时
func SetTimeOut(req *HttpRequest.Request, d time.Duration) *HttpRequest.Request {
	req.SetTimeout(d) //default 30s
	return req
}
