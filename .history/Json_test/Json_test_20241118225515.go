package main

import (
	"encoding/json"
	"fmt"
)

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

	test_p_json, err := json.Marshal(&p)

	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
	}

	fmt.Println("序列化完成：", string(test_p_json))

	return string(test_p_json)
}

// 从json解码
func FromJson(data string) Person {
	// Your code here

	TEST_DATA_byte := []byte(data)
	// TEST_DATA_byte := []byte(TEST_DATA)
	var person_test Person
	if json.Unmarshal(TEST_DATA_byte, &person_test) == nil {
		fmt.Println("json.Unmarshal 解码结果: ", person_test.name, person_test.age)
	}

	return Person{}
}

func main() {

	p_me := &Person{
		name: "WUTONK",
		age:  18,
	}

	p_me.ToJson()
	FromJson(TEST_DATA)

}
