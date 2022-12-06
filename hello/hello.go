package hello

import "fmt"

func Test() {
	fmt.Println("Hello, World!")
	var bdunk Dunk
	var wdunk Dunk
	bdunk = Blackdunk{"zy", 23, "male"}
	wdunk = Whitedunk{"zy", 22, "female"}
	fmt.Println(wdunk)
	fmt.Println(bdunk)
	fmt.Println("happy")
}

type Dunk interface {
	speak()
	run()
}

type Blackdunk struct {
	name string
	age  int
	sex  string
}

type Whitedunk struct {
	name string
	age  int
	sex  string
}

func (b Blackdunk) speak() {
	fmt.Println(b.name)
}

func (b Blackdunk) run() {
	fmt.Println(b.name)
}

func (w Whitedunk) speak() {
	fmt.Println(w.name)
}

func (w Whitedunk) run() {
	fmt.Println(w.name)
}
