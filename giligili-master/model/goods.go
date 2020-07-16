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
	Price      float64
	GoodsPrice float32
	State      uint8
	Unite      string
	Image      string
	Stock      uint64
	Sales      string
	Status     uint8
	GoodsID      uint
	GoodsTypeID  uint
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
	GoodsSKU	GoodsSKU
}

// GoodsType 商品类型
type GoodsType struct {
	gorm.Model
	GoodsId    uint16
	CreateTime string
	UpdateTime string
	State      uint8
	Name       string
	Logo       string
	Image      string
	GoodsSKU	GoodsSKU
}
