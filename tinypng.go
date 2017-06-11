package go_tinypng

import (
	"bytes"
	"encoding/json"
	"path/filepath"

	"io/ioutil"

	"net/http"

	"errors"
)

const (
	compressingUrl = "https://api.tinify.com/shrink"
)

type TinyPngResponse struct {
	Input struct {
		Size int64  `json:"size"`
		Type string `json:"type"`
	} `json:"input"`
	Output struct {
		Size   int64   `json:"size"`
		Type   string  `json:"type"`
		Width  int64   `json:"width"`
		Height int64   `json:"height"`
		Ratio  float32 `json:"ratio"`
		Url    string  `json:"url"`
	} `json:"output"`
}

type TinyPng struct {
	email  string
	apiKey string
}

func NewTinyPng(email string, apiKey string) (*TinyPng, error) {
	if len(email) == 0 {
		return nil, errors.New("email is required")
	}

	if len(apiKey) == 0 {
		return nil, errors.New("apiKey is required")
	}

	t := new(TinyPng)
	t.email = email
	t.apiKey = apiKey

	return t, nil
}

/**
 * @param 	path		图片的路径
 * @return 	ret		压缩后的数据
 * @return 	error	错误
 */
func (t *TinyPng) CompressFile(path string) (ret []byte, err error) {
	path, err = filepath.Abs(path)
	if err != nil {
		return
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	ret, err = t.compressBytes(data)
	return
}

/**
 * @param 	abs 		图片的URL链接
 * @return 	ret		压缩后的数据
 * @return 	error	错误
 */
func (t *TinyPng) CompressUrl(url string) (ret []byte, err error) {
	response, err := http.DefaultClient.Get(url)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	ret, err = t.compressBytes(data)
	return
}

/**
 * @param 	abs 		二进制格式的图片
 * @return 	ret		压缩后的数据
 * @return 	error	错误
 */
func (t *TinyPng) CompressBytes(data []byte) (ret []byte, err error) {
	return t.compressBytes(data)
}

func (t *TinyPng) compressBytes(data []byte) (ret []byte, err error) {
	// Create request
	req, err := http.NewRequest(http.MethodPost, compressingUrl, nil)
	if err != nil {
		return
	}

	//  Authentication request
	req.SetBasicAuth(t.email, t.apiKey)

	// Write picture data to request
	req.Body = ioutil.NopCloser(bytes.NewReader(data))

	// Send request
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	// Parse response
	jsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	jsonStruct := new(TinyPngResponse)
	err = json.Unmarshal(jsonData, jsonStruct)
	if err != nil {
		return
	}

	if len(jsonStruct.Output.Url) == 0 {
		err = errors.New("image url is empty")
		return
	}

	response, err = http.DefaultClient.Get(jsonStruct.Output.Url)
	if err != nil {
		return
	}

	ret, err = ioutil.ReadAll(response.Body)

	return
}
