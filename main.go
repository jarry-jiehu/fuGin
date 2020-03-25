package main

import (
	"yxauth/routers"
	"yxauth/utils"
)

func main() {
	r := routers.SetupRouter()

	log.Info("-------logrus test --------")

	r.Run(":8080")
}
