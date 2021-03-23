package goHttp

import (
	"errors"
)

func HttpGet(httpHost string, httpPort string, httpData map[string]interface{}, contextType string) (string, error) {
	// Send map[string]interface{}{}
	resp, err := httpReq.Get(httpHost+":"+httpPort, httpData)
	if err != nil {
		return "", errors.New("request fail")
	}
	defer resp.Close()

	body, result := resp.Body()
	if result != nil {
		return "", errors.New("response fail")
	}

	return string(body), nil
}

func HttpPost(httpHost string, httpPort string, httpData map[string]interface{}) (string, error) {
	// Send map[string]interface{}{}
	resp, err := httpReq.Get(httpHost+":"+httpPort, httpData)
	if err != nil {
		return "", errors.New("request fail")
	}
	defer resp.Close()

	body, result := resp.Body()
	if result != nil {
		return "", errors.New("response fail")
	}

	return string(body), nil
}
