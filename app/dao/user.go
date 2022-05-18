package dao

import (
	"fmt"
	"netdisk-back/app/pojo"
	"netdisk-back/app/util"
)

func Login(loginName string, passWord string) (pojo.User, error) {
	var user pojo.User
	db, err := util.GetDBConn()
	if err != nil {
		return user, err
	}
	db.Where("user_name=? and pass_word=?", loginName, passWord).First(&user)
	return user, nil
}
func AddUser(user pojo.User) (pojo.User, error) {
	db, err := util.GetDBConn()
	db = db.Create(&user)

	if db.RowsAffected > 0 {
		return user, err
	}
	return user, err
}
func GetUsers() ([]pojo.User, error) {
	db, err := util.GetDBConn()
	var users []pojo.User
	db.Find(&users)
	return users, err
}
func GetStoreUsers() ([]pojo.User, error) {
	db, err := util.GetDBConn()
	var users []pojo.User
	/*db.Find(&users)
	var store pojo.Store
	db.Model(&users).Association("Store").Find(store)*/
	sql := "SELECT * FROM `user` LEFT JOIN file_store ON `user`.id =  file_store.user_id "
	db.Debug().Raw(sql).Scan(&users)
	fmt.Println("一对一", users)
	/*var store pojo.Store
	db.Debug().Preload("User").Find(&store)
	fmt.Println(store)*/

	return nil, err
}
func GetUserById(uid int) (pojo.User, error) {
	db, err := util.GetDBConn()
	var user pojo.User
	db.Where("id=?", uid).First(&user)
	return user, err
}
func GetUserLike(key string) ([]pojo.User, error) {
	db, err := util.GetDBConn()
	var users []pojo.User
	db.Where("concat(id,user_name,login_name,phone_number,email) like ?", "%"+key+"%").Find(&users)
	return users, err
}

func DelUserById(uid int) (bool, error) {
	db, err := util.GetDBConn()
	var user pojo.User
	db.Where("id=?", uid).Delete(&user)

	if db.RowsAffected > 0 {
		return true, err
	}
	return false, err
}
func UpUserById(user pojo.User) (bool, error) {
	db, err := util.GetDBConn()
	db = db.Save(user)

	if db.RowsAffected > 0 {
		return true, err
	}
	return false, err

}
