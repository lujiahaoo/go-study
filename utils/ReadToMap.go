package utils

import (
	"encoding/json"
	"io/ioutil"
)

func ENV(key string) interface{} {
	b, err := ioutil.ReadFile("./.env")
	if err != nil {
		panic("读取配置文件出错1")
	}

	result := make(map[string]interface{})
	if err := json.Unmarshal([]byte(string(b)), &result); err != nil {
		panic("转换json对象失败")
	}

	return result[key]
}
