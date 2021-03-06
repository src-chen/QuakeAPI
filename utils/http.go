package utils

import (
	"QuakeAPI/log"
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type HttpInterface interface {
	DoGet(url string, data map[string]string, headers map[string]string) []byte
	DoPost(url string, data map[string]string, headers map[string]string) []byte
}

type HttpClient struct {
}

func (h *HttpClient) DoGet(url string, data map[string]string, headers map[string]string) []byte {
	return doRequest("GET", url, data, headers)
}

func (h *HttpClient) DoPost(url string, data map[string]string, headers map[string]string) []byte {
	return doRequest("POST", url, data, headers)
}

func doRequest(
	method string,
	url string,
	data map[string]string,
	headers map[string]string) []byte {
	client := http.Client{}
	var req *http.Request
	var err error
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if data != nil {
		bytesData, _ := json.Marshal(data)
		req, err = http.NewRequest(method, url, bytes.NewBuffer(bytesData))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		log.Log("http error:"+err.Error(), log.ERROR)
		return nil
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Log("read error:"+err.Error(), log.ERROR)
			return nil
		}
		return body
	}
	return nil
}
