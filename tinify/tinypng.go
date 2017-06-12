package Tinify

import "errors"

//
//import (
//	"bytes"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"path/filepath"
//)
//
//// 图片压缩的API地址
//const compressingUrl = "https://api.tinify.com/shrink"
//
const VERSION = "1.0"

var (
	key    string
	client *Client
)

func SetKey(set_key string) {
	key = set_key
}

func GetClient() *Client {
	if len(key) == 0 {
		panic(errors.New("Provide an API key with Tinify.setKey(key string)"))
	}

	if client == nil {
		c, err := NewClient(key)
		if err != nil {
			panic(errors.New("Provide an API key with Tinify.setKey(key string)"))
		}
		client = c
	}
	return client
}

//
//var ApiKey string
//
//type CompressResponse struct {
//	Input struct {
//		Size int64  `json:"size"`
//		Type string `json:"type"`
//	} `json:"input"`
//	Output struct {
//		Size   int64   `json:"size"`
//		Type   string  `json:"type"`
//		Width  int64   `json:"width"`
//		Height int64   `json:"height"`
//		Ratio  float32 `json:"ratio"`
//		Url    string  `json:"url"`
//	} `json:"output"`
//}
//
//const (
//	ResizeMethodScale = "scale"
//	ResizeMethodFit   = "fit"
//	ResizeMethodCover = "cover"
//)
//
//type ResizeMethod string
//
//type ResizeOption struct {
//	Method ResizeMethod `json:"method"`
//	Width  int64        `json:"width"`
//	Height int64        `json:"height"`
//}
//
//type resizeRequest struct {
//	Resize ResizeOption `json:"resize"`
//}
//
//type Tinify struct {
//	// 原图
//	originImageBuf []byte
//	OriginImageUrl string
//
//	// 压缩后的图
//	compressedImageBuf []byte
//	compressedImageUrl string
//	compressedResponse *CompressResponse
//
//	// 改变大小后的图片
//	resizedImageBuf []byte
//	resizedOption   *ResizeOption
//}
//
//func (t *Tinify) ToBuffer() (buf []byte, err error) {
//	if len(t.compressedImageBuf) > 0 {
//		buf = t.compressedImageBuf
//		return
//	}
//
//	if len(t.compressedImageUrl) == 0 {
//		err = t.compressToUrl(t.originImageBuf)
//		if err != nil {
//			return
//		}
//	}
//
//	response, err := http.DefaultClient.Get(t.compressedImageUrl)
//	if err != nil {
//		return
//	}
//
//	buf, err = ioutil.ReadAll(response.Body)
//	if err != nil {
//		return
//	}
//
//	t.compressedImageBuf = buf
//	return
//}
//
//func (t *Tinify) ToFile(path string) (err error) {
//	if len(path) == 0 {
//		err = errors.New("path is required")
//		return
//	}
//
//	path, err = filepath.Abs(path)
//	if err != nil {
//		return
//	}
//
//	buf, err := t.ToBuffer()
//	if err != nil {
//		return
//	}
//
//	err = ioutil.WriteFile(path, buf, os.ModePerm)
//	return
//}
//
//func (t *Tinify) Resize(option *ResizeOption) (err error) {
//	if option == nil {
//		err = errors.New("option is required")
//		return
//	}
//
//	// 校验ApiKey
//	if validApiKey() == false {
//		err = errors.New("ApiKey is require, please code 'Tinify.SetKey()'")
//		return
//	}
//
//	if len(t.compressedImageUrl) == 0 {
//		err = t.compressToUrl(t.originImageBuf)
//		if err != nil {
//			return
//		}
//	}
//
//	// Create request
//	req, err := http.NewRequest(http.MethodGet, t.compressedImageUrl, nil)
//	if err != nil {
//		return
//	}
//
//	// Authentication request
//	req.SetBasicAuth("api", ApiKey)
//
//	// Write picture data to request
//	resizeReq := new(resizeRequest)
//	resizeReq.Resize = *option
//	fmt.Printf("\n%+v\n", resizeReq)
//	data, err := json.Marshal(resizeReq)
//	if err != nil {
//		return
//	}
//	req.Body = ioutil.NopCloser(bytes.NewReader(data))
//
//	fmt.Printf("\n%+v\n", req)
//
//	// Send request
//	response, err := http.DefaultClient.Do(req)
//
//	if err != nil {
//		return
//	}
//
//	// Parse response
//	jsonData, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		return
//	}
//
//	t.resizedImageBuf = jsonData
//	t.resizedOption = option
//
//	return
//}
//
//func (t *Tinify) compressToUrl(data []byte) (err error) {
//
//	// 校验ApiKey
//	if validApiKey() == false {
//		err = errors.New("ApiKey is require, please code 'Tinify.SetKey()'")
//		return
//	}
//
//	if len(t.originImageBuf) == 0 {
//		err = errors.New("Origin image buffer is empty, please code 'Tinify.FromFile(), Tinify.FromUrl() or Tinify.FromBuffer()'")
//		return
//	}
//
//	// Create request
//	req, err := http.NewRequest(http.MethodPost, compressingUrl, nil)
//	if err != nil {
//		return
//	}
//
//	// Authentication request
//	req.SetBasicAuth("api", ApiKey)
//
//	// Write picture data to request
//	req.Body = ioutil.NopCloser(bytes.NewReader(data))
//
//	// Send request
//	response, err := http.DefaultClient.Do(req)
//
//	if err != nil {
//		return
//	}
//
//	// Parse response
//	jsonData, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		return
//	}
//
//	jsonStruct := new(CompressResponse)
//	err = json.Unmarshal(jsonData, jsonStruct)
//	if err != nil {
//		return
//	}
//
//	if len(jsonStruct.Output.Url) == 0 {
//		err = errors.New("image url is empty")
//		return
//	}
//
//	t.compressedResponse = jsonStruct
//
//	t.compressedImageUrl = jsonStruct.Output.Url
//
//	return
//}
//
//func SetKey(key string) error {
//	if len(key) == 0 {
//		return errors.New("key is required")
//	}
//	ApiKey = key
//	return nil
//}
//
//func validApiKey() bool {
//	return len(ApiKey) > 0
//}
//
//func FromFile(path string) (t *Tinify, err error) {
//	// 转为绝对路径
//	path, err = filepath.Abs(path)
//	if err != nil {
//		return
//	}
//
//	data, err := ioutil.ReadFile(path)
//	if err != nil {
//		return
//	}
//
//	t = new(Tinify)
//	t.originImageBuf = data
//	return
//}
//
//func FromUrl(url string) (t *Tinify, err error) {
//	// 获取远程图片
//	response, err := http.DefaultClient.Get(url)
//	if err != nil {
//		return
//	}
//
//	// 读取请求结果
//	data, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		return
//	}
//
//	t = new(Tinify)
//	t.originImageBuf = data
//	return
//}
//
//func FromBuffer(data []byte) (t *Tinify, err error) {
//	t = new(Tinify)
//	t.originImageBuf = data
//	return
//}
