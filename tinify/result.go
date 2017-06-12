package Tinify

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Result struct {
	data []byte
	*ResultMeta
}

func NewResult(meta http.Header, data []byte) *Result {
	r := new(Result)
	r.ResultMeta = NewResultMeta(meta)
	r.data = data
	return r
}

func (r *Result) Data() []byte {
	return r.data
}

func (r *Result) ToBuffer() []byte {
	return r.Data()
}

func (r *Result) ToFile(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, r.data, os.ModePerm)
	return err
}

func (r *Result) Size() int64 {
	s := r.meta["Content-Length"]
	if len(s) == 0 {
		return 0
	}

	size, _ := strconv.Atoi(s[0])
	return int64(size)
}

func (r *Result) MediaType() string {
	arr := r.meta["Content-Type"]
	if len(arr) == 0 {
		return ""
	}
	return arr[0]
}

func (r *Result) ContentType() string {
	return r.MediaType()
}
