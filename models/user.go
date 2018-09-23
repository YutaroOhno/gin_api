package models
import (
    "log"
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
)
var engine *xorm.Engine

func init() {
    var err error
    engine, err = xorm.NewEngine("mysql", "root:@/bookshelf")
    if err != nil {
        log.Fatal(err)
    }
}

type User struct {
    ID       int    `json:"id" xorm:"id pk autoincr"`
    NickName string `json:"nickname" xorm:"'nickname'"`
}
