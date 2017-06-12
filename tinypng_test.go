package main

import (
	"testing"

	"io/ioutil"

	"github.com/astaxie/beego/logs"
	"github.com/gwpp/go-tinypng/tinify"
)

const Key = "rcPZm3Zrg_1DbjYtV6AXM_-53Jg9wuWB"

func TestFromFile(t *testing.T) {
	Tinify.SetKey(Key)
	source, err := Tinify.FromFile("./test.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	err = source.ToFile("./test_output/CompressFromFile.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	logs.Info("压缩成功")
}

func TestFromBuffer(t *testing.T) {
	Tinify.SetKey(Key)

	buf, err := ioutil.ReadFile("./test.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	source, err := Tinify.FromBuffer(buf)
	if err != nil {
		t.Error(err)
		return
	}

	err = source.ToFile("./test_output/CompressFromBuffer.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	logs.Info("压缩成功")
}

func TestFromUrl(t *testing.T) {
	Tinify.SetKey(Key)
	url := "http://pic.tugou.com/realcase/1481255483_7311782.jpg"
	source, err := Tinify.FromUrl(url)
	if err != nil {
		t.Error(err)
		return
	}
	err = source.ToFile("./test_output/CompressFromUrl.jpg")
	if err != nil {
		t.Error(err)
		return
	}
	logs.Info("压缩成功")
}
