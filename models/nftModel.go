package Models

import (
	"fmt"
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)


func (n *Nft) TableName() string {
	return "nft"
}



func GetList(page int, pageNum int) (list []User) {

	o := orm.NewOrm()

	num, err := o.Raw("SELECT id, address, created_at FROM user WHERE id >= ?", 1).QueryRows(&list)
	if err == nil {
	    fmt.Println("user nums: ", num)
	}

	return
}

// Address     string `orm:"column(address);description(用户地址)" json:"address"`
// TokenId     string `orm:"column(tokenId);description(nft tokenId)" json:"tokenId"`
// Name        string `orm:"column(name);size(21);description(nft name)" json:"name"`
// Status      uint8  `orm:"column(status);description(NFT状态)" json:"status"`
// CreatedAt   time.Time `orm:"auto_now_add;type(datetime);description(创建时间)" json:"created_at"`


func InsertNft(address string, tokenId string, name string ) {
    o := orm.NewOrm()

    nft := new(Nft)
    nft.Address = address
    nft.TokenId = tokenId
    nft.Name = name
    nft.Status = 1

    o.Insert(nft)

}

func GetNFTId(address string, tokenId string, name string ) {
    o := orm.NewOrm()

    nft := new(Nft)
    nft.Address = address
    nft.TokenId = tokenId
    nft.Name = name
    nft.Status = 1

    o.Insert(nft)
    
}










