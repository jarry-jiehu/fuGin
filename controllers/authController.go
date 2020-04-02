package controllers

import (
	"github.com/gin-gonic/gin"
	"fuGin/utils"
)

type AuthController struct {
	BaseController
}

type AuthInfo struct {
	UserName string
	Password string
}

func (this *AuthController) AuthLogIn(c *gin.Context) {
	this.c = c

	name := c.DefaultPostForm("userName", "")
	password := c.DefaultPostForm("password", "")

	in := AuthInfo{
		UserName: name,
		Password: password,
	}

	log.Info("name: %s, pwd: %s", name, password)

	this.response(200, "success", in)
}

func (this *AuthController) AuthSignIn(c *gin.Context) {

}
