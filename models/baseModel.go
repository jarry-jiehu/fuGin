package models

import (
	"fmt"
	"fuGin/utils"
	"github.com/Unknwon/goconfig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var cfg *goconfig.ConfigFile

func GetCfg() *goconfig.ConfigFile {
	if cfg == nil {
		var err error
		cfg, err = goconfig.LoadConfigFile("conf/app.ini")
		if err != nil {
			log.Fatalf("Read config file failed at path %s,%s", "conf/app.ini", err)
		}
	}
	return cfg
}

func init() {
	// 创建数据库及redis，包括后面使用的kafka，rabbitmq，zk等
	// 数据库使用gorm框架
	initGorm()
}

func initGorm() {
	mysqlMap, err := GetCfg().GetSection("mysql")
	if err != nil {
		log.Error("Get conf error!")
	}
	var host string = mysqlMap["db.mysql.host"]
	var username string = mysqlMap["db.mysql.username"]
	var password string = mysqlMap["db.mysql.password"]
	var dbName string = mysqlMap["db.mysql.dbname"]

	dbstr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, dbName)

	log.Infof("---- dbcmd: %s", dbstr)
	_, err = gorm.Open("mysql", dbstr)

	if err != nil {
		log.Info(err)
		fmt.Println(err)
		return
	} else {
		log.Info("connection succedssed")
		fmt.Println("connection succedssed")
	}
}
