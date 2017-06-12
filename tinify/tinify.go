package Tinify

import "errors"

const VERSION = "1.0"

var (
	key    string
	client *Client
)

func SetKey(set_key string) {
	key = set_key
}

func GetClient() *Client {
	if len(key) == 0 {
		panic(errors.New("Provide an API key with Tinify.setKey(key string)"))
	}

	if client == nil {
		c, err := NewClient(key)
		if err != nil {
			panic(errors.New("Provide an API key with Tinify.setKey(key string)"))
		}
		client = c
	}
	return client
}
