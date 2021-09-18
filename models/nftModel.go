package Models

import (
	"fmt"
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)


func (n *Nft) TableName() string {
	return "nft"
}



func GetList(skip int, take int) (list []Nft) {

	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("id", "address", "name", "tokenId", "image", "status").
		From("nft").
		Where("status = 1").
		OrderBy("id").Asc().
    	Limit(take).Offset(skip)

    sql := qb.String()


	num, err := o.Raw(sql).QueryRows(&list)
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


func InsertNft(address, tokenId, name, image string) {
    o := orm.NewOrm()

    nft := new(Nft)
    nft.Address = address
    nft.TokenId = tokenId
    nft.Name = name
    nft.Image = image
    nft.Status = 1

    o.Insert(nft)

}

func GetNFTId(address, tokenId, name string ) {
    o := orm.NewOrm()

    nft := new(Nft)
    nft.Address = address
    nft.TokenId = tokenId
    nft.Name = name
    nft.Status = 1

    o.Insert(nft)
    
}










