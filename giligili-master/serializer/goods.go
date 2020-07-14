package serializer

import "giligili/model"

// Movie 视频序列化器
type Goods struct {
	GoodsId    uint16 `json:"goodsId"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	Status     uint8  `json:"status"`
	Name       string `json:"name"`
	Content    string `json:"content"`
}

// BuildGoods 序列化电影
func BuildGoods(item model.Goods) Goods {
	return Goods{
		GoodsId:    item.GoodsId,
		CreateTime: item.CreateTime,
		UpdateTime: item.UpdateTime,
		Status:     item.Status,
		Name:       item.Name,
		Content:    item.Content,
	}
}

// BuildManyGoods 序列化电影列表
func BuildManyGoods(items []model.Goods) (manyGoods []Goods) {
	for _, item := range items {
		goods := BuildGoods(item)
		manyGoods = append(manyGoods, goods)
	}
	return manyGoods
}
