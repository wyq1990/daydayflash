package api

import (
	"fmt"
	"giligili/service/fresh"

	"github.com/gin-gonic/gin"
)

func GetName(name string) {
	fmt.Printf(name)
}

// ListMovie 视频列表接口
func ListGoods(c *gin.Context) {
	service := fresh.ListGoodsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
