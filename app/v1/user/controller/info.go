package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/tuuz/RET"
)

func InfoController(route *gin.RouterGroup) {
	route.Use(BaseController.CommonController())
	route.Use(BaseController.LoginedController())
	route.Any("get", info_get)
}

func info_get(c *gin.Context) {
	uid := c.GetHeader("uid")
	if user, err := UserModel.Api_findById(uid); err != nil {
		RET.Fail(c, 500, nil, nil)
	} else {
		RET.Success(c, 0, user, nil)
	}
}
