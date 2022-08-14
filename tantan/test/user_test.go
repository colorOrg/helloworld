package test

import (
	"fmt"
	"hello/tantan/bean"
	dao "hello/tantan/dao/fun"
	"hello/tantan/service"
	"testing"
	"time"
)

func TestGetUser(t *testing.T)  {

	userRes,err:=service.GetUsers()
	if (err!=nil){
		fmt.Println(err)
	}
	fmt.Print(userRes)
}

func TestAddUser(t *testing.T)  {

	var q = (*bean.User)(nil)
	fmt.Println(q)

	var p = bean.NewUser("11")
	fmt.Println(p.Type)



}

func TestConvertUser(t *testing.T)  {
	array:=[]string{"暂不透露"}
	v := checkUserExtensionsArrayEmpty(array)
	fmt.Println(v)
}

func checkUserExtensionsArrayEmpty(array []string) bool {
	if len(array) == 0 {
		return true
	}

	for i := range array {
		if len(array[i]) != 0 {
			return false
		}
	}
	return true
}

func Test1(t *testing.T) {
	r := dao.NewExpertRepository()
	expert := &dao.Expert{
		UserID:            111,
		ExpertName:        "name",
		ExpertAvatar:      "http://www.baidu.com",
		ExpertDescription: "hahah",
		PhoneNumber:       "18810383",
		ExpertStatus:      "default",
		CreatedTime:       time.Time{},
		UpdatedTime:       time.Time{},
	}
	err:=r.InsertExpert(nil, expert)
	fmt.Sprintf("err%d%s", expert.ID,err)
}




