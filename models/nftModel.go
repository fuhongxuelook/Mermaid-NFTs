package Models

import (
	"fmt"
    "strconv"
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)


func (n *Nft) TableName() string {
	return "nft"
}



func GetList(skip int, take int, query, address string) (list []Nft) {
    list = []Nft{}

	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("id", "address", "name", "tokenId", "image", "status").
		From("nft").
		Where("status > 1")

    if address != "" && address != "0"  {
        qb.And("address = '" + address + "'")
    }

    if query != "" {
        qb.And("(tokenId='" + query + "' or name like '%" + query + "%')") 
    }



	qb.OrderBy("id").Asc().
    	Limit(take).Offset(skip)

    sql := qb.String()


	num, err := o.Raw(sql).QueryRows(&list)
	if err == nil {
	    fmt.Println("user nums: ", num)
	}

	return
}



func GetImageByTokenId(tokenId string) (image string) {

    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("image").
        From("nft").
        Where("status > 1").
        And("tokenId='" + tokenId + "'")

    sql := qb.String()


    err := o.Raw(sql).QueryRow(&image)
    if err == nil {
        fmt.Println("image: ", image)
    }


    return 
}

func GetListNum(skip int, take int, query, address string) (num int64) {

    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("count(*) as num").
        From("nft").
        Where("status > 1")

    if address != "" && address != "0"  {
        qb.And("address = '" + address + "'")
    }

    if query != "" {
        qb.And("(tokenId='" + query + "' or name like '%" + query + "%')") 
    }

   

    sql := qb.String()


    err := o.Raw(sql).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }


    return num
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

func GetNFTId() (string) {
    tokenId := -1
    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("TokenId").
        From("nft").
        OrderBy("TokenId").
        Desc()

        //Where("status = 1")

    sql := qb.String()


    err := o.Raw(sql).QueryRow(&tokenId)
    if err == nil {
        fmt.Println("tokenId is: ", tokenId)
    }

    if tokenId == -1 {
        tokenId = 0
    } else {
        tokenId += 1
    }


    return strconv.Itoa(tokenId)
    
}

func ChangeNftTokenIdStatus(tokenId, status string) {
    o := orm.NewOrm()

    res, err := o.Raw("UPDATE nft SET status = ? WHERE tokenId = ?", status, tokenId).Exec()
    if err == nil {
        num, _ := res.RowsAffected()
        fmt.Println("mysql row affected nums: ", num)
    }

    return

}










