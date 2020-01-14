/**
  create by yy on 2020/1/9
*/

package main

import (
	"fmt"
	"go_crawler/libs"
	"net/http"
)

var (
	MarkMap map[int]int
)

type responseData struct {
	Id           int    `json:"id"`
	OperatorCode string `json:"operator_code"`
}

func main() {
	MarkMap = make(map[int]int, 1)
	http.HandleFunc("/callback", func(writer http.ResponseWriter, request *http.Request) {
		unlockId := request.URL.Query()

		fmt.Println(unlockId)

		libs.WriteError(writer, "success")
	})

	http.HandleFunc("/get_operator", func(writer http.ResponseWriter, request *http.Request) {

		libs.WriteSuccess(writer, "success")
	})

	http.ListenAndServe("127.0.0.1:8089", nil)
}

// func SendHourData(campID int, reqtType string) {
// 	URL := ""
// 	sendDataURL := ""
// 	if reqtType == PROD {
// 		URL = "http://www.preadmin.com/get/newest/click/hour?camp_id=" + strconv.Itoa(campID)
// 		sendDataURL = "http://www.preadmin.com/insert/click/hour/data"
// 	} else {
// 		URL = "http://127.0.0.1:8080/get/newest/click/hour?camp_id=" + strconv.Itoa(campID)
// 		sendDataURL = "http://127.0.0.1:8080/insert/click/hour/data"
// 	}
// 	newestDate, _ := httplib.Get(URL).String()
// 	logs.Info("111111111", newestDate)
// 	if newestDate != "FALSE" && campID != 0 {
// 		hourClick := new(HourClick)
// 		totalHourData, err := hourClick.GetCampClick(campID, newestDate)
//
// 		if err == nil {
// 			var sendDataList []HourClick
// 			for i, oneData := range *totalHourData {
// 				sendDataList = append(sendDataList, oneData)
// 				if i%100 == 0 { // 每次最多发送100条数据
// 					sendData(sendDataURL, sendDataList) // 发送数据
//
// 					sendDataList = []HourClick{} // 清空 list数据
// 				}
// 			}
//
// 			sendData(sendDataURL, sendDataList) // 发送剩余的数据
//
// 		}
// 	}
// }
