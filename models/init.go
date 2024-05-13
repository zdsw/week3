package models

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db     *gorm.DB
	Client *elastic.Client
	err    error
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/week3?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}

	err = Db.AutoMigrate(&TeachingInfoTable{}, &UploadTable{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}

	Client, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9201"), elastic.SetSniff(false))
}
