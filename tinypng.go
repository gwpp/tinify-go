package main

import (
	"net/http"

	"io/ioutil"

	"bytes"

	"fmt"

	"time"

	"github.com/astaxie/beego/logs"
)

const (
	CompressingUrl = "https://api.tinify.com/shrink"

	Email  = "ganwenpwng_dev@163.com"
	ApiKey = "rcPZm3Zrg_1DbjYtV6AXM_-53Jg9wuWB"
)

func init() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
}

func main() {
	// 创建Request
	req, err := http.NewRequest(http.MethodPost, CompressingUrl, nil)
	if err != nil {
		logs.Error(err)
		return
	}

	// 将鉴权信息写入Request
	req.SetBasicAuth(Email, ApiKey)

	// 将图片以二进制的形式写入Request
	data, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		logs.Error(err)
		return
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(data))

	// 发起请求
	fmt.Print("正在进行网络请求，由于Tinypng是国外网站，国内访问速度可能会比较慢，请稍耐心等待")
	go func() {
		for {
			fmt.Print(".")
			time.Sleep(1 * time.Second)
		}
	}()
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.Error(err)
		return
	}

	// 解析请求
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		logs.Error(err)
		return
	}

	fmt.Print("\n")
	fmt.Println("请求完毕，结果如下：")
	fmt.Println(string(data))
}
