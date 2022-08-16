package controller

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"hello/tantan/bean"
	"hello/tantan/convert"
	"hello/tantan/service"
	"io/ioutil"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users,err := service.GetUsers()
	if err != nil {
		res := bean.Response{500, err.Error(), ""}
		fmt.Fprint(w, res)
		return
	}

	msg, _ := json.Marshal(users)
	res := bean.Response{200, "", string(msg)}
	fmt.Fprint(w, res)
}


func AddUser(w http.ResponseWriter, r *http.Request) {
	var req map[string] interface{}
	body,_ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	var user =bean.NewUser("")
	err := mapstructure.Decode(req, &user)
	if err!=nil {
		res := bean.Response{400, "输入参数不正确", ""}
		fmt.Fprint(w, res)
		return
	}
	err = service.AddUsers(user)
	userRes,err := convert.ConvertUserToUserRes(user)
	if err!=nil {
		res := bean.Response{400, "新增用户失败", ""}
		fmt.Fprint(w, res)
		return
	}


	res := bean.Response{200, "", userRes.String()}
	fmt.Fprint(w, res)

}
