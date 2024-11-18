package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	P_ME := &Person{
		name: "WUTONK",
		age:  18,
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
	name string
	age  int64
}

// 转为Json
func (p *Person) ToJson() string {

	TestPJson, err := json.Marshal(&p)

	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
	}

	fmt.Println("序列化完成：", string(TestPJson))

	return string(TestPJson)
}

// 从json解码
func FromJson(data string) Person {
	// Your code here

	TEST_DATA_byte := []byte(data)
	var PersonTest Person //接受解码结果
	if json.Unmarshal(TEST_DATA_byte, &PersonTest) == nil {
		fmt.Println("json.Unmarshal 解码结果: ", PersonTest.name, PersonTest.age)
	}

	return Person{}
}
