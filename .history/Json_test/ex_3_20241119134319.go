package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	P_ME := &Person{
		Name: "WUTONK",
		Age:  18,
		City: "Canton",
	}

	PMeString := P_ME.ToJson()
	fmt.Println("序列化完成：", PMeString)

	PersonTest := FromJson(TEST_DATA)
	fmt.Println("json.Unmarshal 解码结果: ", PersonTest)

}

const (
	TEST_DATA  = `{"name": "John", "age": 30, "city": "New York"}`
	TEST_DATA2 = `{"name": "Person", "age": -328}`
)

type Person struct {
	// Your code here
	Name string
	Age  int64
	City string
}

// 序列化为Json
func (p *Person) ToJson() string {

	TestPJson, err := json.Marshal(&p)

	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
		return ""
	}

	return string(TestPJson)
}

// 从json反序列化
func FromJson(data string) Person {
	// Your code here

	Test_Date_Byte := []byte(data)
	var PersonTest Person //接受解码结果
	if json.Unmarshal(Test_Date_Byte, &PersonTest) != nil {
		println("出现错误")
	}

	return PersonTest
}