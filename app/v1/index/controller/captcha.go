package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/AossGoSdk"
	"github.com/tobycroft/Calc"
	"main.go/common/BaseController"
	"main.go/config/app_conf"
	"main.go/tuuz/RET"
)

func CaptchaController(route *gin.RouterGroup) {
	route.Use(BaseController.CommonController())
	route.Any("/get", captcha_get)
}

func captcha_get(c *gin.Context) {
	var captcha AossGoSdk.Captcha
	captcha.Token = app_conf.Project
	oid := Calc.GenerateOrderId()
	img, err := captcha.Math(oid)
	if err != nil {
		RET.Fail(c, 500, nil, err)
		return
	}
	RET.Success(c, 0, map[string]interface{}{
		"ident": oid,
		"img":   img,
	}, nil)
}
