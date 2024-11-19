package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type add_int struct {
	// 接受样例：`{"a": 12, "b": 18}`
	A_INT int `json:"a"`
	B_INT int `json:"b"`
}

type add_string struct {
	// 接受样例：`{ "a": "12", "b": "18" }`
	A_STR string `json:"a"`
	B_STR string `json:"b"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/currentTime", func(c *gin.Context) {
		RFC3339 := "2006-01-02T15:04:05Z07:00" //go中统一使用go诞生时间作为时间模版
		TimeRes := time.Now().Format(RFC3339)
		c.JSON(200, gin.H{"message": TimeRes})
	})

	r.POST("/add", func(c *gin.Context) {
		//接受数据的变量
		json := add_int{}
		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.BindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		num_a := json.add_int_json_a
		num_b := json.add_int_json_b

		num_sub := num_a + num_b

		c.JSON(200, gin.H{
			"元数据a": json.add_int_json_a,
			"元数据b": json.add_int_json_b,
			"相加结果": num_sub,
		})

	})

	r.POST("/addString", func(c *gin.Context) {
		json := add_string{}

		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.BindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//接受参数
		string_a := json.A_String
		string_b := json.B_String

		String_Sum := string_a + string_b

		c.JSON(200, gin.H{
			"元数据a": json.A_String,
			"元数据b": json.B_String,
			"相加结果": String_Sum,
		})
	})

	r.Run("127.0.0.1:11451")

}