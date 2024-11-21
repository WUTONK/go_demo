package main

import (
	"fmt"
)

// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
type Utau interface {
	Say()
}

type Miku struct {
}

func (miku Miku) say() {
	fmt.Println("Miku Don't talk to the British")
}

type Teto struct {
}

func (teto Teto) say() {
	fmt.Println("ユ!")
}

func main() {
	// utau := Utau{}
	var utau Utau

	utau = new(Miku)
	utau.Say()

}
