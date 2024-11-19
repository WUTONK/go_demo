package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	personMe := &Person{
		Name: "WUTONK",
		Age:  18,
		City: "Canton",
	}

	pMeJson := personMe.ToJson()
	fmt.Println("序列化完成：", pMeJson)

	personTest := FromJson(TEST_DATA)
	fmt.Println("json.Unmarshal 解码结果: ", personTest)

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

	toJsonPerson, err := json.Marshal(&p)

	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
		return ""
	}

	return string(toJsonPerson)
}

// 从json反序列化
func FromJson(data string) Person {
	// Your code here

	testDateByte := []byte(data)
	var personTest Person //接受解码结果
	if json.Unmarshal(testDateByte, &personTest) != nil {
		println("出现错误")
	}

	return personTest
}
