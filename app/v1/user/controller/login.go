package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/AossGoSdk"
	"github.com/tobycroft/Calc"
	"main.go/app/v1/user/model/UserModel"
	"main.go/common/BaseController"
	"main.go/common/BaseModel/TokenModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Base64"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func LoginController(route *gin.RouterGroup) {
	route.Use(BaseController.CommonController())

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
	mail, ok := Input.PostLength("mail", 6, 16, c, true)
	if !ok {
		return
	}
	password, ok := Input.PostLength("password", 6, 16, c, false)
	if !ok {
		return
	}
	passmd5 := Calc.Md5(password)
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
	dataUser, err := UserModel.Api_findByEmail(mail)
	if err != nil {
		RET.Fail(c, 500, nil, err)
		return
	}
	token := Calc.GenerateToken()
	id := int64(0)
	if dataUser == nil {
		id, err = UserModel.Api_insert(mail, mail, passmd5)
		if err != nil {
			RET.Fail(c, 500, nil, err)
			return
		}
	} else {
		id = Calc.Any2Int64(dataUser["id"])
		if dataUser["password"] != passmd5 {
			RET.Fail(c, 400, nil, "密码错误，账号已被注册，你可以通过对应的邮箱再次找回")
			return
		}
	}
	err = TokenModel.Api_insert(id, token, "default")
	if err != nil {
		RET.Fail(c, 500, nil, err)
		return
	}
	RET.Success(c, 0, map[string]any{
		"token": token,
		"uid":   id,
	}, "注册成功")
}
