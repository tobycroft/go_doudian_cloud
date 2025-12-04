package controller

import (
	"github.com/gin-gonic/gin"
)

func CaptchaController(route *gin.RouterGroup) {
	route.Any("/captcha", captcha_get)
}

func captcha_get(c *gin.Context) {

	c.String(200, "ok")
}
