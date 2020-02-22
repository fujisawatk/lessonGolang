package main

import (
	"fmt"
)

type UserNotFound struct {
	Username string
}

// カスタムエラー
// →Printlnエラー表記を自前で設定できる。
func (e *UserNotFound) Error() string{
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func main() {
	// エラー内容が違う場所で発生した時、ポインタ出ないと同一のエラーとして扱われるため、
	// &を記述する。
	e1 := &UserNotFound{"mike"}
	e2 := &UserNotFound{"mike"}
	// &をつけないと同一エラーとしてtrueが返される
	fmt.Println(e1 == e2)
	if err := myFunc(); err != nil {
		fmt.Println(err)
		// エラーの種類によって処理条件をつける時、
		// 別々のエラーとして認識されなければならない。
		if err == e1 {

		}

		if err == e2 {

		}
	}
}