# go-tinypng
---

## 前言
所有大前端（Web、Android、iOS）的同学应该都会遇到这样一个需求——压缩图片资源体积。Web端是为了加快页面加载速度，App端是为了减小apk、ipa的体积，但不管怎么说目的是一样的——压缩。

查找过相关资料的同学估计都知道一个神奇的网站——[Tinypng](https://tinypng.com)，我们可以很简单方便的在线压缩图片资源，而且Tinypng的压缩比例还是很可观的。

Tinypng很好的解决了我们压缩图片的需求，而且提供了Online、API、PS-Plugin等方式供我们使用，API方面官方提供了许多语言的SDK支持，但遗憾的是并没有golang的，所以今天带大家使用golang来实现一下Tinypng的压缩。

## Code

- 核心代码

```
package main

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"github.com/astaxie/beego/logs"
)

const (
	CompressingUrl = "https://api.tinify.com/shrink"

	// Email和ApiKey替换成自己的
	Email  = "example@163.com"
	ApiKey = "rcPxm3Zrg_1DbjYtV6AXM_-53Jg9wuWB"
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

	logs.Info(string(data))
}
```
- 输出

```
{
  "input":{
    "size":322199,
    "type":"image/jpeg"
  },
  "output":{
    "size":141938,
    "type":"image/jpeg",
    "width":2880,
    "height":1800,
    "ratio":0.4405,
    "url":"https://api.tinify.com/output/fg5ibhadc16kbf4h.jpg"
  }
}
```
- 解析
  - 关于Email、ApiKey
    使用Tinypng的Api服务时需要注册他们平台的账号，注册地址：https://tinypng.com/developers 。注册需要提供邮箱和用户名，用户名随意填写，邮箱就是我们代码中的Email。注册后平台会给你提供的邮箱发一封邮件，里面有个链接，点击去就可以看到ApiKey。
  - 关于API相关信息
    详情见官方HTTP文档：https://tinypng.com/developers/reference
  - 关于日志
    demo中的log使用的是beego的log模块，这里也感谢开源库beego
  - 关于结果
    demo最后只是把Response的body输出了，但是从输出结果的json也不难看出其中包含着图片的所有信息，我们通过 `output.url`即可获取压缩后的图片。

## 写在最后
这个demo虽然实现了功能，但确实过于简单，之后有时间我会尝试将Tinypng所涉及到的功能多封装在一起，以此弥补官方不提供golang SDK的遗憾。
