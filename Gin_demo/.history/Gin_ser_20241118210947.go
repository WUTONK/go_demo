package Gin_demo

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/currentTime", func(c *gin.Context) {
		TimeRes := time.Now().Unix()
		c.JSON(200, gin.H{"message": TimeRes})
	})

	r.POST("/add" ,func(c *gin.Context) {
		
	})


	r.Run("127.0.0.1:11451")

	func Json_test{

	};
	
};


