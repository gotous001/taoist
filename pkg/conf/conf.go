package conf

import (
	"encoding/json"
	"io"
)

var _ methods = &conf{}

var Conf *conf

type methods interface {
	//Conf() *conf
	Get(key string) interface{}
}

type conf struct {
	KeyMap map[string]interface{}
}

func newConf(reader io.Reader) *conf {
	keyMap := make(map[string]interface{})

	decoder := json.NewDecoder(reader)
	_ = decoder.Decode(&keyMap)

	return &conf{
		keyMap,
	}
}

func (c *conf) Get(key string) interface{} {
	resp := c.KeyMap[key]
	return resp
}
