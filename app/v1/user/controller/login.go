package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/AossGoSdk"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/user/model/TokenModel"
	"main.go/app/v1/user/model/UserModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Base64"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func LoginController(route *gin.RouterGroup) {
	route.Use(cors.Default())

	route.Any("captcha", login_captcha)
	route.Any("auto", login_auto)
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

func login_auto(c *gin.Context) {
	var captcha AossGoSdk.Captcha
	captcha.Token = app_conf.Project
	username, ok := Input.PostLength("username", 1, 16, c, true)
	if !ok {
		return
	}
	password, ok := Input.PostLength("password", 6, 16, c, false)
	if !ok {
		return
	}
	mail, ok := Input.PostLength("mail", 6, 16, c, true)
	if !ok {
		return
	}
	if !app_conf.TestMode {
		ident, ok := Input.Post("ident", c, false)
		if !ok {
			return
		}
		code, ok := Input.Post("code", c, false)
		if !ok {
			return
		}
		err := captcha.CheckInTime(ident, code, 600)
		if err != nil {
			RET.Fail(c, 500, nil, err)
			return
		}
	}
	if dataMail, err := UserModel.Api_findByEmail(mail); err != nil {
		RET.Fail(c, 500, nil, err)
		return
	} else if dataMail != nil {
		RET.Fail(c, 401, nil, "邮箱已被注册，你可以申请找回或重设密码")
		return
	}
	if id, err := UserModel.Api_insert(username, mail, password); err != nil {
		RET.Fail(c, 500, nil, err)
	} else {
		token := Calc.GenerateToken()
		err = TokenModel.Api_insert(id, token, "default")
		if err != nil {
			RET.Fail(c, 500, nil, err)
			return
		}
		RET.Success(c, 0, map[string]any{
			"token": token,
			"id":    id,
		}, "注册成功")
	}
}
