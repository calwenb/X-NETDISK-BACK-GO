package pojo

import "time"

type User struct {
	Id           int       `gorm:"primaryKey,AUTO_INCREMENT" json:"userId"`
	UserName     string    `json:"userName"`
	LoginName    string    `json:"loginName"`
	PassWord     string    `json:"passWord"`
	UserType     int       `gorm:"default:2" json:"userType"`
	PhoneNumber  string    `json:"phoneNumber"`
	Email        string    `json:"email"`
	Avatar       string    `gorm:"default:/*" json:"avatar"`
	RegisterTime time.Time `gorm:"autoCreateTime" json:"registerTime"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}
