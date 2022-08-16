package dao

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"hello/tantan/bean"
)

func GetRelationshipByUId(db *pg.DB, uId uint64) ([]bean.Relationship, error){
	var rs []bean.Relationship
	err := db.Model(&rs).Where("u_id_f = ?", uId).WhereOr("u_id_t = ?" ,uId).Select()
	return rs,err
}


func GetRelationshipByDirect(db *pg.DB, uIdF uint64, uIdT uint64, t uint8) ([]bean.Relationship, error){
	var rs []bean.Relationship
	var err error
	if t==1 {
		err = db.Model(&rs).Where("u_id_f = ?", uIdF).Where("u_id_t = ?", uIdT).Select()
	} else {
		err = db.Model(&rs).WhereOrGroup(func(q *orm.Query) (query *orm.Query, err error) {
			q=q.Where("u_id_f = ?", uIdF).Where("u_id_t = ?", uIdT)
			return q,nil
		}).WhereOrGroup(func(q *orm.Query) (query *orm.Query, err error) {
			q=q.Where("u_id_f = ?", uIdT).Where("u_id_t = ?", uIdF)
			return q,nil
		}).Select()
	}

	return rs,err
}

func AddRelationship(db * pg.DB, rs *bean.Relationship)  (orm.Result, error){
	res,err:=db.Model(rs).Returning("id").Insert()
	return res,err
}

func UpdateRelationship(db * pg.DB, rs *bean.Relationship)  (orm.Result, error){
	res,err:=db.Model(rs).Set("state_f=?", rs.StateF).Set("state_t=?", rs.StateT).Where("id=?id").Update();
	return res,err
}



