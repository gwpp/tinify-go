package Tinify

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego/logs"
)

const API_ENDPOINT = "https://api.tinify.com"
const RETRY_COUNT = 1
const RETRY_DELAY = 500

type Client struct {
	options map[string]interface{}
	key     string
}

func NewClient(key string) (c *Client, err error) {
	c = new(Client)
	c.key = key
	return
}

// method: http.MethodPost、http.MethodGet
func (c *Client) Request(method string, url string, body interface{}) (response *http.Response, err error) {
	if strings.HasPrefix(url, "https") == false {
		url = API_ENDPOINT + url
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	logs.Info(body)

	switch body.(type) {
	case []byte:
		if len(body.([]byte)) > 0 {
			req.Body = ioutil.NopCloser(bytes.NewReader(body.([]byte)))
		}
	case map[string]interface{}:
		if len(body.(map[string]interface{})) > 0 {
			body2, err2 := json.Marshal(body)
			if err2 != nil {
				err = err2
				return
			}
			req.Body = ioutil.NopCloser(bytes.NewReader(body2))
		}
		req.Header["Content-Type"] = []string{"application/json"}
		logs.Info("%+v", req.Header)
	}

	req.SetBasicAuth("api", c.key)

	logs.Info("%+v", req)

	response, err = http.DefaultClient.Do(req)
	return
}
