package libs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Reply struct {
	Writer http.ResponseWriter `json:"-"`
	Code   int                 `json:"code"`
	Data   interface{}         `json:"data"`
	Msg    interface{}         `json:"msg"`
}

func GetReply(response http.ResponseWriter, code int, data ...interface{}) (reply *Reply) {

	var result interface{}

	// 设置返回数据格式
	response.Header().Set("Content-Type", "application/json")

	// 设置成功状态码
	response.WriteHeader(http.StatusOK)

	if len(data) < 1 {
		result = ""
	} else {
		result = data[0]
	}

	reply = &Reply{
		Code:   code,
		Writer: response,
	}

	if code != 0 {
		reply.Msg = result
		reply.Data = ""
	} else {
		reply.Msg = ""
		reply.Data = result
	}

	return
}

func (r *Reply) Write() {

	var (
		data []byte
		err  error
	)

	if data, err = json.Marshal(r); err != nil {
		fmt.Println(NewReportError(err).Error())
	}

	// 设置输出
	if _, err = r.Writer.Write(data); err != nil {
		fmt.Println(NewReportError(err).Error())
	}

}

func GetSuccess(response http.ResponseWriter, data ...interface{}) *Reply {
	return GetReply(response, 0, data...)
}

func GetError(response http.ResponseWriter, data ...interface{}) *Reply {
	return GetReply(response, -1, data...)
}

func WriteSuccess(response http.ResponseWriter, data ...interface{}) {
	GetSuccess(response, data...).Write()
}

func WriteError(response http.ResponseWriter, data ...interface{}) {
	GetError(response, data...).Write()
}

// custom code and write message
func WriteCustomCode(response http.ResponseWriter, code int, data ...interface{}) {
	GetReply(response, code, data...).Write()
}

type LoginReply struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}
