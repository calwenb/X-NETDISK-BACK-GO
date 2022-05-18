package service

import (
	"github.com/pkg/errors"
	"netdisk-back/app/dao"
	"netdisk-back/app/pojo"
	"strconv"
)

func Login(userName string, passWord string) (bool, pojo.User) {
	user, _ := dao.Login(userName, passWord)
	if user == (pojo.User{}) {
		return false, user
	}
	return true, user
}

func Register(user pojo.User) (bool, pojo.User) {
	user, _ = dao.AddUser(user)
	if user.Id == 0 {
		return false, user
	}
	return true, user
}

func GetUsers() []pojo.User {
	users, _ := dao.GetUsers()
	return users
}
func GetStoreUsers() []pojo.User {
	users, _ := dao.GetStoreUsers()
	return users
}

func GetUserLike(key string) []pojo.User {
	users, _ := dao.GetUserLike(key)
	return users
}

func GetUserById(uid string) pojo.User {
	uidI, _ := strconv.Atoi(uid)
	user, _ := dao.GetUserById(uidI)
	return user
}

func DelUserById(uid string) bool {
	uidI, err := strconv.Atoi(uid)
	if err != nil {
		errors.New("转换错误")
		return false
	}
	rs, _ := dao.DelUserById(uidI)
	return rs
}

func UpUserById(user pojo.User) bool {
	rs, _ := dao.UpUserById(user)
	return rs
}
