package main

import (
	"GateWay/balancer"
	"GateWay/config"
	"GateWay/handle"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	config, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("read config file error : %s", err)
	}

	err = config.Validation()
	if err != nil {
		log.Fatalf("verify config error : %s", err)
	}

	// Location中pattern标准化
	config.Normalization()

	// Pattern to balancer init
	err = balancer.InitP2b(config.Location)
	if err != nil {
		log.Fatalf("p2b error : %s", err)
	}

	r := gin.Default()
	for _, l := range config.Location {
		switch l.Method {
		case "GET":
			r.GET(l.Pattern, handle.GenHandlerFunc(l.RemovePrefix))
		case "POST":
			r.POST(l.Pattern, handle.GenHandlerFunc(l.RemovePrefix))
		case "PUT":
			r.PUT(l.Pattern, handle.GenHandlerFunc(l.RemovePrefix))
		case "DELETE":
			r.DELETE(l.Pattern, handle.GenHandlerFunc(l.RemovePrefix))
		}
	}
	// 跨域预检
	//r.OPTIONS("/*action", func(context *gin.Context) {
	//
	//})

	r.Run(":" + strconv.Itoa(config.Port))
}
