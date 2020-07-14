package model

import "github.com/jinzhu/gorm"

// GoodsSKU 商品sku
type GoodsSKU struct {
	gorm.Model
	GoodsId    uint16
	CreateTime string
	UpdateTime string
	Name       string
	Content    string
	Price      uint64
	Unite      string
	Image      string
	Stock      uint64
	Sales      string
	Status     uint8
	Goods      Goods
	GoodsType  GoodsType
}

// Goods 商品spu
type Goods struct {
	gorm.Model
	GoodsId    uint16
	CreateTime string
	UpdateTime string
	Status     uint8
	Name       string
	Content    string
}

// GoodsType 商品类型
type GoodsType struct {
	gorm.Model
	GoodsId    uint16
	CreateTime string
	UpdateTime string
	Name       string
	Logo       string
	Image      string
}
