package pool

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var DBPool *gorm.DB

func InitDB() {
	oldDsn := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		"root",
		"1234",
		"127.0.0.1",
		"3307",
		"account_system",
	)
	oldDb, err := gorm.Open("mysql", oldDsn)
	if err != nil {
		fmt.Printf("mysql connection failure, error: (%v)", err.Error())
		return
	}
	oldDb.DB().SetMaxIdleConns(10)  // 设置连接池
	oldDb.DB().SetMaxOpenConns(100) // 设置与数据库建立连接的最大数目
	oldDb.DB().SetConnMaxLifetime(time.Second * 7)
	DBPool = oldDb
}
