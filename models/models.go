package Models

import (
	"time"
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
    Id          int    `orm:"column(id)" json:"id"`
    Address     string `orm:"column(address);description(地址)" json:"address"`
    Status      uint8  `orm:"column(status);description(用户状态)" json:"status"`
    TokenId     string `orm:"column(tokenId);description(new user nft tokenId)" json:"tokenId"`
    CreatedAt   time.Time `orm:"auto_now_add;type(datetime);description(创建时间)" json:"created_at"`
}

type Nft struct {
    Id          int    `orm:"column(id)" json:"id"`
    Address     string `orm:"column(address);description(用户地址)" json:"address"`
    TokenId     string `orm:"column(tokenId);description(nft tokenId)" json:"tokenId"`
    Name        string `orm:"column(name);size(21);description(nft name)" json:"name"`
    Image       string `orm:"column(image);size(100);description(image name)" json:"image"`
    Status      uint8  `orm:"column(status);description(NFT状态)" json:"status"`
    CreatedAt   time.Time `orm:"auto_now_add;type(datetime);description(创建时间)" json:"created_at"`
}


func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    // set default database
    orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/mermaidnft?charset=utf8")

    orm.Debug = true
    // // // register model
    // orm.RegisterModel(new(Stu))
    orm.RegisterModel(new(User), new(Nft))

    orm.RunSyncdb("default", false, true)


    // // // create table
    // orm.RunSyncdb("mermaidnft", false, true)
}