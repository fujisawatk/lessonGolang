package main

import "fmt"

func goroutine1(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
		c <- sum
	}
	// 処理が終了したら、channelも閉じる
	// 閉じてないとmain関数のfor文で取り出し続ける処理が続くため、エラーが起きる
	close(c)
}

func main(){
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine1(s, c)
	// channelで受信されるたびに、iが出力
	for i := range c{
		fmt.Println(i) // 1 3 6 10 15
	}
}