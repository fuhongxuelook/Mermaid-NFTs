package Models

import (
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)


func (u *User) TableName() string {
	return "user"
}

func InsertUser(address string, tokenId string) {
	o := orm.NewOrm()

    user := new(User)
    user.Address = address
    user.Status = 1
    user.TokenId = tokenId

    o.Insert(user)
}













