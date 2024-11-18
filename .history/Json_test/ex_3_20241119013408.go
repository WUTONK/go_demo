package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	P_ME := &Person{
		Name: "WUTONK",
		Age:  18,
	}

	P_ME.ToJson()
	FromJson(TEST_DATA)

}

const (
	TEST_DATA  = `{"name": "John", "age": 30, "city": "New York"}`
	TEST_DATA2 = `{"name": "Person", "age": -328}`
)

type Person struct {
	// Your code here
	Name string
	Age  int64
}

// 转为Json
func (p *Person) ToJson() string {

	TestPJson, err := json.Marshal(&p)

	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
		return ""
	}

	fmt.Println("序列化完成：", string(TestPJson))

	return string(TestPJson)
}

// 从json解码
func FromJson(data string) Person {
	// Your code here

	Test_Date_Byte := []byte(data)
	var PersonTest Person //接受解码结果
	if json.Unmarshal(Test_Date_Byte, &PersonTest) == nil {
		fmt.Println("json.Unmarshal 解码结果: ", PersonTest.name, PersonTest.age)
	}

	return Person{}
}
