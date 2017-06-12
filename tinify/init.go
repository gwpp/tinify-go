package Tinify

import "github.com/astaxie/beego/logs"

func init() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
}
