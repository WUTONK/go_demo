package Gin_demo

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	http.Lis
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
