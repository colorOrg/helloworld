package service

import (
	"hello/tantan/bean"
	"hello/tantan/convert"
	"hello/tantan/dao"
)

func GetUsers() (*[]bean.UserResponse, error) {
	users,err := dao.GetUsers(dao.Db)
	userRes,err := convert.ConvertUsersToUserRes(&users)
	return userRes,err
}

func AddUsers(user *bean.User)  (error) {
	_, err := dao.AddUser(dao.Db, user)
	return err
}





