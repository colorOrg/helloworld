package dao

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"hello/tantan/bean"
)




func GetUsers(db *pg.DB) ([]bean.User, error){
	var users []bean.User
	err:=db.Model(&users).Select()
	return users,err
}

func AddUser(db * pg.DB, user *bean.User)  (orm.Result, error){
	res,err:=db.Model(user).Returning("id").Insert()
	return res,err
}




