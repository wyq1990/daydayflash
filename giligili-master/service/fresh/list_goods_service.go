package fresh

import (
	"giligili/model"
	"giligili/serializer"
)

// ListGoodsService 视频列表服务
type ListGoodsService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListGoodsService) List() serializer.Response {
	goods := []model.Goods{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 15
	}

	if err := model.DB.Model(model.Goods{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&goods).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildManyGoods(goods), uint(total))
}
