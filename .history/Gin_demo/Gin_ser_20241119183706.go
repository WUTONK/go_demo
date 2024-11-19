package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type AddInt struct {
	// 接受样例：`{"a": 12, "b": 18}`
	A_INT int `json:"a"`
	B_INT int `json:"b"`
}

type AddString struct {
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
		rfc3339 := "2006-01-02T15:04:05Z07:00" //go中统一使用go诞生时间作为时间模版
		timeRes := time.Now().Format(rfc3339)
		c.String(http.StatusOK, timeRes)
	})

	r.POST("/add", func(c *gin.Context) {
		//接受数据的变量
		json := AddInt{}
		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.BindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		numSum := json.A_INT + json.B_INT

		c.JSON(200, gin.H{"sum": numSum})

	})

	r.POST("/addString", func(c *gin.Context) {
		json := AddString{}

		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.BindJSON(&json); err != nil {
			// 返回错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//接受参数
		numA, err := strconv.Atoi(json.A_STR)
		if err != nil {
			println("转码错误:", err)
		}
		numB, err := strconv.Atoi(json.B_STR)
		if err != nil {
			println("转码错误:", err)
		}

		stringSum := numA + numB

		c.JSON(200, gin.H{"sum": stringSum})
	})

	r.Run("127.0.0.1:11451")

}
