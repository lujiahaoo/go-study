package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

//返回响应状态码及body
func ResponseWithJson(w http.ResponseWriter, payload interface{}, statusCode ...int) {
	if len(statusCode) > 1 {
		panic("too many arguments of ResponseWithJson()")
	}
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")

	//因为Go不支持有默认值的可选参数函数，所以用了这种骚方法
	code := http.StatusOK
	if len(statusCode) == 1 {
		code = statusCode[0]
	}

	w.WriteHeader(code)
	w.Write(response)
}

// MarshalJson 把对象以json格式放到response中
// func (resp *Resp) MarshalJson(w http.ResponseWriter) error {
// 	data, err := json.Marshal(resp)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Fprintln(w, string(data)) //data是[]byte类型，转化成string类型便于查看
// 	return nil
// }

// UnMarshalJson 从request中取出对象
func (resp *Response) UnMarshalJson(req *http.Request) (interface{}, error) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return data, err
	}
	var result interface{}
	json.Unmarshal(data, result)
	return result, nil
}
