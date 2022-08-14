package convert

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"hello/tantan/bean"
)

func ConvertUserToUserRes(user *bean.User)  (*bean.UserResponse, error){
	temp, err:= json.Marshal(user)
	var m map[string]interface{}
	err = json.Unmarshal([]byte(temp), &m)
	switch m["type"].(float64){
		case 1:
			m["type"]="user"
	}
	uRes := bean.UserResponse{}
	mapstructure.Decode(m, &uRes)
	return &uRes, err
}

func ConvertUsersToUserRes(users *[]bean.User)  (*[]bean.UserResponse, error){
	var err error
	usersRes := make([]bean.UserResponse, len(*users))
	for i,user := range *users {
		userC,err := ConvertUserToUserRes(&user)
		if (err!=nil){
			panic(err)
		}
		usersRes[i] = *userC
	}
	return &usersRes,err
}




func ConvertRsToRsRes(rs *bean.Relationship, uId uint64) (*bean.RelationshipResponse, error){
	temp, err:= json.Marshal(rs)
	var m map[string]interface{}
	err = json.Unmarshal([]byte(temp), &m)
	m["type"]="relationship"
	switch m["state"].(float64){
		case 1:
			m["state"]="liked"
		case 2:
			m["state"]="disliked"
		case 3:
			m["state"]="matched"
	}
	if rs.UIdF == uId  {
		m["userId"]=rs.UIdT
	} else {
		m["userId"]=rs.UIdF
	}
	rsRes := bean.RelationshipResponse{}
	mapstructure.Decode(m, &rsRes)
	return &rsRes, err
}

func ConvertRssToRsRes(rss *[]bean.Relationship, uId uint64)  (*[]bean.RelationshipResponse, error) {
	var err error
	rssRes := make([]bean.RelationshipResponse, len(*rss))
	for i,rs := range *rss {
		rsC,err := ConvertRsToRsRes(&rs, uId)
		if err!=nil {
			panic(err)
		}
		rssRes[i] = *rsC
	}
	return &rssRes,err
}

