package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type add_int struct {
	// 接受样例：`{"a": 12, "b": 18}`
	add_int_json_a int64 `json:"a"`
	add_int_json_b int64 `json:"b"`
}

type add_string struct {
	// 接受样例：`{ "a": "12", "b": "18" }`
	add_string_json_a string `json:"a"`
	add_string_json_b string `json:"b"`
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

		// 读取请求体
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取请求体"})
			return
		}
		fmt.Printf("Received JSON: %s\n", body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 重新设置请求体

		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.BindJSON(&json); err != nil {
			// 返回错误信息
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

		// 读取请求体
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取请求体"})
			return
		}
		fmt.Printf("Received JSON: %s\n", body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 重新设置请求体

		//将request的body中的数据，自动按照json格式解析到结构体
		if err := c.BindJSON(&json); err != nil {
			// 返回错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//接受参数
		string_a := json.add_string_json_a
		string_b := json.add_string_json_b

		string_sum := string_a + string_b

		c.JSON(200, gin.H{
			"元数据a": json.add_string_json_a,
			"元数据b": json.add_string_json_b,
			"相加结果": string_sum,
		})
	})

	r.Run("127.0.0.1:11451")
}