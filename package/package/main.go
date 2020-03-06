package main

import (
	// 自作パッケージを指定するときは、$GOPASTHからのディレクトリを記述する。
	"fmt"
	"lessonGolang/package/package/mylib"
	"lessonGolang/package/package/mylib/under"
)

func main() {
	// スライス定義
	s := []int{1, 2, 3, 4, 5}
	// パッケージmylibの関数Averageに変数sを引数で渡す。
	// 処理されたものを出力
	fmt.Println(mylib.Average(s))

	// パッケージmylibの関数Sayを呼び出す。
	mylib.Say()

	// パッケージunderの関数Helloを呼び出す。
	under.Hello()
}
