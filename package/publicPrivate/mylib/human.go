package mylib

import "fmt"

// グローバル変数定義
var Public string = "Public"
var private string = "private"

// 構造体Person定義
type Person struct {
	Name string
	Age  int
}

// 関数定義
func Say() {
	fmt.Println("Human!")
}

/* 上記を定義する場合
public：頭文字を大文字にすると他のパッケージでも扱える。
private：頭文字を小文字にすると他のパッケージでは扱えない。
*/
