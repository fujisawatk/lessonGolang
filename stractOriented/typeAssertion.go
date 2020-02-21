package main

import "fmt"

// interfaceを使うと、どの型でも引数で扱える
// タイプアサーション
// →変数.(型)で値の型を明示する
func do(i interface{}){
	// switch type文
	// →switch文のみtype表記ででタイプアサーションできる。
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main(){
	do(10)
	do("Mike")
	do(true)

	// タイプコンバージョン（キャスト）
	// →intをfloat64で明示することで型を変更できる
	// ※タイプアサーションといているが記述方法が違うため注意
	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)
}