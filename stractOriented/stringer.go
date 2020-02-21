package main

import "fmt"

type Person struct {
	Name string
	Age int
}

// Stringer
// →Printlnで出力を操作したい時、Stringメソッドで指定できる。
func (p Person) String() string {
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}