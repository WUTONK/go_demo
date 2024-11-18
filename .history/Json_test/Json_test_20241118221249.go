package main

const (
	TEST_DATA  = `{"name": "John", "age": 30, "city": "New York"}`
	TEST_DATA2 = `{"name": "Person", "age": -328}`
)

type Person struct {
	// Your code here
	name string
	age  int64
}

func (p *Person) ToJson() string {

	return ""
}

func FromJson(data string) Person {
	// Your code here
	return Person{}
}

func main() {
	// Your code here

	test_to_json := Person{
		name: "WUTONK",
		age:  18,
	}
}
