package go_tinypng

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	Email  = "ganwenpwng_dev@163.com"
	ApiKey = "rcPZm3Zrg_1DbjYtV6AXM_-53Jg9wuWB"
)

func TestTinyPng_CompressFile(t *testing.T) {
	t.Log("Begin test TestTinyPng_CompressFile")
	tin, err := NewTinyPng(Email, ApiKey)
	if err != nil {
		t.Error(err)
		return
	}

	ret, err := tin.CompressFile("./test.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	err = ioutil.WriteFile("./test_output/TestTinyPng_CompressFile.jpg", ret, os.ModePerm)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("TestTinyPng_CompressFile: %v", ret)
}

func TestTinyPng_CompressBytes(t *testing.T) {
	t.Log("Begin test TestTinyPng_CompressBytes")
	tin, err := NewTinyPng(Email, ApiKey)
	if err != nil {
		t.Error(err)
		return
	}

	data, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	ret, err := tin.CompressBytes(data)
	if err != nil {
		t.Error(err)
		return
	}
	err = ioutil.WriteFile("./test_output/TestTinyPng_CompressBytes.jpg", ret, os.ModePerm)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("TestTinyPng_CompressBytes: %v", ret)
}

func TestTinyPng_CompressUrl(t *testing.T) {
	t.Log("Begin test TestTinyPng_CompressUrl")
	tin, err := NewTinyPng(Email, ApiKey)
	if err != nil {
		t.Error(err)
		return
	}

	ret, err := tin.CompressUrl("http://pic.tugou.com/realcase/1481255483_7311782.jpg")
	if err != nil {
		t.Error(err)
		return
	}

	err = ioutil.WriteFile("./test_output/TestTinyPng_CompressUrl.jpg", ret, os.ModePerm)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("TestTinyPng_CompressUrl: %v", ret)
}
