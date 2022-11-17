package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// a test sever 8082
func main() {
	r := gin.Default()
	r.GET("/v1/api/cc", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "i am 8082:/v1/api/cc",
		})
	})

	r.Run(":8082")
}
