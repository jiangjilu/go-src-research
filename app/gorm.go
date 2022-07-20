package app

import (
	"fmt"
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
	Name     string
	Products []Product `gorm:"many2many:product_class;"`
}

var db *gorm.DB

func Gorm() {
	initDb()
	db.AutoMigrate(&Product{})
	createRow()

	// 查询关联
	productA := Product{
		Model: gorm.Model{
			ID: 1,
		},
	}

	// 关联计数
	count := db.Model(&productA).Association("Classes").Count()
	println(count)

	// 删除关联
	db.Model(&productA).Association("Classes").Clear()
	count2 := db.Model(&productA).Association("Classes").Count()
	println(count2)

	// 更新关联
	productUpdate := Product{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "40GP",
		Classes: []Class{
			{
				Model: gorm.Model{
					ID: 1,
				},
			},
			{
				Model: gorm.Model{
					ID: 2,
				},
			},
		},
	}

	// 替换关联
	//db.Model(&productA).Association("Classes").Replace(productUpdate.Classes)

	db.Save(productUpdate)

	var product Product
	db.Preload("Classes").Find(&product, 1)

	fmt.Println(product)
}

func initDb() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/go_src_research?charset=utf8mb4&parseTime=True&loc=Local"
	dbn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = dbn
}

func createRow() {
	fcl := Product{
		Model: gorm.Model{},
		Name:  "40HQ",
		Classes: []Class{
			{
				Model: gorm.Model{},
				Name:  "Ocean",
			},
		},
	}
	db.Create(&fcl)
}
