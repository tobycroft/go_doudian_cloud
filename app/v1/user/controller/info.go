package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
)

func InfoController(route *gin.RouterGroup) {
	route.Use(BaseController.CommonController())
	route.Use(BaseController.LoginedController())
	route.Any("/info", info_get)
}

func info_get(c *gin.Context) {

}
