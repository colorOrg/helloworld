package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"hello/tantan/bean"
	"hello/tantan/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetRelationship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uId :=vars["user_id"]
	num,err :=strconv.ParseUint(uId, 10, 64)
	if (err!=nil){
		fmt.Fprint(w, "输入参数有误")
		return
	}
	rs,err := service.GetRelationships(num)
	if err != nil {
		res := bean.Response{500, err.Error(), ""}
		fmt.Fprint(w, res)
		return
	}
	msg, _ := json.Marshal(rs)
	res := bean.Response{200, "", string(msg)}
	fmt.Fprint(w, res)
}


func PutRelationship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uIdF :=vars["user_id"]
	uIdT :=vars["other_user_id"]
	var req map[string] interface{}
	body,_ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	uIdFNum, uIdTNum, stateNum, err := validateRelationshipParam(uIdF, uIdT, req)
	if (err != nil) {
		fmt.Fprint(w, err.Error())
		return
	}
	rss,err := service.PutRelationship(uIdFNum, uIdTNum, stateNum)
	if err != nil {
		res := bean.Response{500, err.Error(), ""}
		fmt.Fprint(w, res)
		return
	}

	msg, _ := json.Marshal(rss)
	res := bean.Response{200, "", string(msg)}
	fmt.Fprint(w, res)

}

func validateRelationshipParam(uIdF string, uIdT string, req map[string] interface{}) (uint64, uint64, uint8, error){

	uIdFNum,err :=strconv.ParseUint(uIdF, 10, 64)
	uIdTNum,err :=strconv.ParseUint(uIdT, 10, 64)
	state :=req["state"].(string)
	var stateNum uint8
	if (state == "liked") {
		stateNum=1
	} else if (state == "disliked") {
		stateNum=2
	} else {
		err= errors.New("state 参数传递错误")
	}
	return uIdFNum, uIdTNum, stateNum, err
}





