package main

import (
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
}

func main() {

}
