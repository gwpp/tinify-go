package Tinify

import (
	"net/http"

	"strconv"
)

type ResultMeta struct {
	meta http.Header
}

func NewResultMeta(meta http.Header) *ResultMeta {
	r := new(ResultMeta)
	r.meta = meta
	return r
}

func (r *ResultMeta) width() int64 {
	w := r.meta["Image-Width"]
	if len(w) == 0 {
		return 0
	}
	width, _ := strconv.Atoi(w[0])

	return int64(width)
}

func (r *ResultMeta) height() int64 {
	h := r.meta["Image-Height"]
	if len(h) == 0 {
		return 0
	}

	height, _ := strconv.Atoi(h[0])
	return int64(height)
}

func (r *ResultMeta) location() string {
	arr := r.meta["Location"]
	if len(arr) == 0 {
		return ""
	}
	return arr[0]
}
