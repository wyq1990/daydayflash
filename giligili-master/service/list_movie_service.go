package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ListMovieService 视频列表服务
type ListMovieService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListMovieService) List() serializer.Response {
	movies := []model.Movie{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 15
	}

	if err := model.DB.Model(model.Movie{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&movies).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildMovies(movies), uint(total))
}
