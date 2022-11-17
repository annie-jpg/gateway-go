package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// a test server 8081
func main() {
	r := gin.Default()
	r.GET("/v1/api/cc", func(context *gin.Context) {
		println(context.Param("action"))
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "i am 8081:/v1/api/cc",
		})
	})

	r.Run(":8081")
}
