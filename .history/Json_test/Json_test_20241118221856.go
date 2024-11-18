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

func (p *Person, to_json_value Person) ToJson() string {

	test_p_json, err := json.Marshal(&to_json_value)

	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
	}

	fmt.Printf("序列化完成：", test_p_json)

	return ""
}

func FromJson(data string) Person {
	// Your code here
	return Person{}
}

func main() {
	// Your code here

	test_p := Person{
		name: "WUTONK",
		age:  18,
	}

	Person.ToJson(Person, test_p)

}
