package main

import (
			"time"
			"fmt"
)


func goroutine1(ch chan string){
	for {
		ch <- "packet from 1"
		time.Sleep(3 * time.Second)
	}
}

func goroutine2(ch chan int){
	// 1秒おきにchへデータを送信
	for {
		ch <- 100
		time.Sleep(1 * time.Second)
	}
}

func main(){
	c1 := make(chan string)
	c2 := make(chan int)
	go goroutine1(c1)
	go goroutine2(c2)

	for {
		// 各channelが受信したものをcaseで分けて出力。
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}