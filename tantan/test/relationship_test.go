package test

import (
	"fmt"
	"hello/tantan/bean"
	"hello/tantan/convert"
	"hello/tantan/service"
	"reflect"
	"testing"
	"time"
)

func TestGetRs(t *testing.T)  {

	rs,err := service.GetRelationships(1)
	if (err!=nil){
		fmt.Println(err)
	}
	fmt.Print(rs)
}

func TestAddRs(t *testing.T)  {
	c :="b"

	m := map[string]HelperMessageABConfig{"a": {}}
	isOk,_:=m[c]
	fmt.Println(isOk)
	fmt.Println(reflect.DeepEqual(isOk,HelperMessageABConfig{}))

}

type HelperMessageABConfig struct {
	ABGroupName string
	ABGroupRules map[string]HelperMessageABRule // AB rules
	UpdatedTime  time.Time
}

type HelperMessageABRule struct {
	RuleName string
	ABRuleInfo map[string]string // AB rule
	UpdatedTime  time.Time
}

func TestConvertRs(t *testing.T)  {
	rs := bean.NewRelationship(1,2,1, 1)
	rs.State=0
	rsRes,err := convert.ConvertRsToRsRes(rs, 1)
	if (err!=nil){
		panic(err)
	}
	fmt.Println(rsRes.String())
}


