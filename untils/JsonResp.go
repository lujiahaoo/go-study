package untils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Resp struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// MarshalJson 把对象以json格式放到response中
func (resp *Resp) MarshalJson(w http.ResponseWriter) error {
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, string(data)) //data是[]byte类型，转化成string类型便于查看
	return nil
}

// UnMarshalJson 从request中取出对象
func (resp *Resp) UnMarshalJson(req *http.Request) (interface{}, error) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return data, err
	}
	var result interface{}
	json.Unmarshal(data, result)
	return result, nil
}
