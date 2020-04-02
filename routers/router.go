package routers

import (
	"github.com/gin-gonic/gin"
	"fuGin/controllers"
)

var ginRouter *gin.Engine

func init() {

	ginRouter = gin.Default()

}

func SetupRouter() *gin.Engine {
	g1 := ginRouter.Group("/auth")

	authCtrl := new(controllers.AuthController)
	g1.POST("/login", authCtrl.AuthLogIn)
	g1.POST("/signin", authCtrl.AuthSignIn)

	return ginRouter
}
