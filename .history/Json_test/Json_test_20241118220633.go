package main

const (
	TEST_DATA  = `{"name": "John", "age": 30, "city": "New York"}`
	TEST_DATA2 = `{"name": "Person", "age": -328}`
)

func main() {
	// Your code here
}

type Person struct {
	// Your code here
	name string
	age  int64
	city string
}

func (p *Person) examples() {
	p.name = "WUTONK"
	p.age = 18
	p.city = "Canton"
}

func (p *Person) ToJson() string {

	return ""
}

func FromJson(data string) Person {
	// Your code here
	return Person{}
}