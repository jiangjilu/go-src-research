package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1个产品属于多个分类
type Product struct {
	gorm.Model
	Name    string
	Classes []Class `gorm:"many2many:product_class;"`
}

// 1个分类包括多个产品
type Class struct {
	gorm.Model
	Name string
}

func Gorm() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/go_src_research?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	_ = db.AutoMigrate(&Product{})

	oceanClass := Class{
		Model: gorm.Model{},
		Name:  "Ocean",
	}

	var classes []Class
	oceanClasses := append(classes, oceanClass)

	fcl := Product{
		Model:   gorm.Model{},
		Name:    "40HQ",
		Classes: oceanClasses,
	}

	db.Create(&fcl)

}
