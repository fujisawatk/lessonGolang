package main

import (
	"fmt"
	"lessonGolang/package/publicPrivate/mylib"
)

func main() {
	mylib.Say()

	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)

	// publicのため読み込める
	fmt.Println(mylib.Public)
	// privateのため読み込めない
	// fmt.Println(mylib.private)
}
