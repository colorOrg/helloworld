package service

import (
	"hello/tantan/bean"
	"hello/tantan/convert"
	"hello/tantan/dao"
)

func GetRelationships(uId uint64) (*[]bean.RelationshipResponse, error) {
	users,err := dao.GetRelationshipByUId(dao.Db, uId)
	CheckRelationshipsState(users)
	usersRes,err := convert.ConvertRssToRsRes(&users,uId)
	return usersRes,err
}


func CheckRelationshipsState(rss []bean.Relationship){
	m := make([]*bean.Relationship, len(rss))
	for k, _ := range rss {
		m[k] = &rss[k]
	}
	for _,rs := range m {
		CheckRelationshipState(rs)
	}
}

func CheckRelationshipState(rs *bean.Relationship){
	if (rs.StateF==1 && rs.StateT==0) || (rs.StateF==0 && rs.StateT==1) {
		rs.State=1
	} else if (rs.StateF==2 || rs.StateT==2) {
		rs.State=2
	} else if (rs.StateF==1 && rs.StateT==1) {
		rs.State=3
	} else {
		rs.State=2
	}
}

func PutRelationship(uIdF uint64, uIdT uint64, state uint8) (*bean.RelationshipResponse, error){
	rss,err := dao.GetRelationshipByDirect(dao.Db, uIdF, uIdT, 2)
	var rs *bean.Relationship
	if (rss==nil){
		// 新增
		rs =bean.NewRelationship(uIdF, uIdT, state, 0)
		_, err = dao.AddRelationship(dao.Db, rs)
	} else {
		// update
		rs = &rss[0]
		if (rs.UIdF == uIdF && rs.UIdT == uIdT) {
			rs.StateF = state
		} else {
			rs.StateT = state
		}
		_, err = dao.UpdateRelationship(dao.Db, rs)
	}
	CheckRelationshipState(rs)
	rsRes,err:=convert.ConvertRsToRsRes(rs, uIdT)
	return rsRes, err

}


