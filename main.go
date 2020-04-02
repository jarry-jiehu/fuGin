package main

import (
	"fuGin/routers"
	"fuGin/utils"
	"fuGin/models"
)

func main() {
	r := routers.SetupRouter()

	log.Info("-------logrus test --------")

	add, _ := models.GetCfg().GetValue("redis", "address")
	name, _ := models.GetCfg().GetValue("sec1", "name")
	test, _ := models.GetCfg().GetValue("", "test")

	redissec, _ := models.GetCfg().GetSection("redis")
	str := redissec["password"]
	log.Infof("-------logrus test -------- %s", str)

	log.Infof("-------logrus test -------- %s", add)
	log.Infof("-------logrus test -------- %s", name)
	log.Infof("-------logrus test -------- %s", test)

	r.Run(":8081")
}