package main

import (
	"net/http"
	"strconv"
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
	// 接受样例：`{ "a": "12", "b": "18" }`
	add_string_json_a string `form:"a" json:"a" uri:"a" xml:"a" binding:"required"`
	add_string_json_b string `form:"b" json:"b" uri:"b" xml:"b" binding:"required"`
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

		//接受参数
		num_a, err := strconv.Atoi(json.add_int_json_a)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		num_b, err := strconv.Atoi(json.add_int_json_b)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		num_sub := num_a + num_b

		c.JSON(200, gin.H{"相加结果": num_sub})

	})

	r.POST("/addString", func(c *gin.Context) {
		var json add_string
		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//接受参数
		string_a := json.add_string_json_a
		string_b := json.add_string_json_b

		string_sum := string_a + string_b

		c.JSON(200, gin.H{"相加结果": string_sum})
	})

	r.Run("127.0.0.1:11451")

}
