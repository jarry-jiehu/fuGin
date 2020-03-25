package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
	c *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (this *BaseController) response(code int, message string, data interface{}) {

	resp := &Response{
		Code: code,
		Msg:  message,
		Data: data,
	}

	this.c.JSON(http.StatusOK, resp)
}
