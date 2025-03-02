package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/conf"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}
