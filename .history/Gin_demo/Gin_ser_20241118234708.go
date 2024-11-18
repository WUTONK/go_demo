package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type add_int struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	// 接受样例：`{"a": 12, "b": 18}`
	add_int_json_a string `form:"a" json:"a" uri:"a" xml:"a" binding:"required"`
	add_int_json_b string `form:"b" json:"b" uri:"b" xml:"b" binding:"required"`
}

type add_string struct {
	add_string_json string `{ "a": "12", "b": "18" }`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/currentTime", func(c *gin.Context) {
		TimeRes := time.Now().Unix()
		c.JSON(200, gin.H{"message": TimeRes})
	})

	r.POST("/add", func(c *gin.Context) {
		//接受数据的变量
		var json add_int
		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		num_a := string(json.add_int_json_a)
	})

	r.Run("127.0.0.1:11451")

}
