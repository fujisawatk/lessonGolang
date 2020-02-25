package main

import "fmt"

func main(){
	// 受信するchannel数を指定
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch)) // 1
	ch <- 200
	fmt.Println(len(ch)) // 2
	// channel数が指定されているので、それ以上入れようとしたらエラー
	// ch <- 300
	// fmt.Println(len(ch))

	// channelから値を取り出す
	// x := <-ch
	// fmt.Println(x)　// 100

	// fmt.Println(len(ch)) // 1
	// 空きができたので、新たに値を入れることができる。
	// ch <- 300
	// fmt.Println(len(ch)) // 2

	// for&range文を使う場合

	// channelの終了宣言。
	close(ch)
  // closeしないとループで取り出そうとするためエラーが起きる。
	for c := range ch{
		fmt.Println(c)
	}
}