package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/AossGoSdk"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/index/model/UserModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Base64"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func LoginController(route *gin.RouterGroup) {
	route.Use(cors.Default())

	route.Any("captcha", login_captcha)
	route.Any("login", login_login)
}

func login_captcha(c *gin.Context) {
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
		"img":   Base64.EncodePng(img),
	}, nil)
}

func login_login(c *gin.Context) {
	var captcha AossGoSdk.Captcha
	captcha.Token = app_conf.Project
	ident, ok := Input.Post("ident", c, false)
	if !ok {
		return
	}
	code, ok := Input.Post("code", c, false)
	if !ok {
		return
	}
	err := captcha.CheckInTime(ident, code, 500)
	if err != nil {
		RET.Fail(c, 500, nil, err)
		return
	}
	username, ok := Input.Post("username", c, false)
	if !ok {
		return
	}
	password, ok := Input.Post("password", c, false)
	if !ok {
		return
	}
	mail, ok := Input.Post("mail", c, true)
	if !ok {
		return
	}

	if err := UserModel.Api_insert(username, mail, password); err != nil {
		RET.Fail(c, 500, nil, err)
	} else {
		RET.Success(c, 0, nil, "注册成功")
	}
}
