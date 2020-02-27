package main

import "fmt"

func goroutine1(s []int, c chan int){
	sum := 0
	// ループ処理（インデックス番号入らないのでアンダーバー）
	for _, v := range s{
		sum += v
	}
	// for文が終わったらsumをcに送る
	c <- sum
}

func goroutine2(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum
}

func main(){
	// スライス宣言
	s := []int{1, 2, 3, 4, 5}
	// channel宣言
	// unbuffer→関数の順番関係なく処理されたものから入れてく。
	c := make(chan int)
	// 関数にs, cを渡す
	go goroutine1(s, c)
	go goroutine2(s, c)
	// cに受信されたら動作（timeなどで待機させなくても勝手に待機してくれる）
	// channelで受信されたものから順に処理されていく。
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}