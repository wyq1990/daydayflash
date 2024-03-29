package api

import (
	// "giligili/api/fresh"
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// ListMovie 视频列表接口
func ListMovie(c *gin.Context) {
	// fresh.GetName("123")

	service := service.ListMovieService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
