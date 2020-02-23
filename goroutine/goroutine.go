package main

import (
	"fmt"
	"sync"
	// "time"
)

func goroutine(s string, wg *sync.WaitGroup){
	// defer wg.Done()
	for i := 0; i < 5; i++{
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	// 上記のdeferで宣言しても良い。
	wg.Done()
}

func normal(s string){
	for i := 0; i < 5; i++{
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	// syncパッケージのWaitGroupメソッドを宣言
	var wg sync.WaitGroup
	// 一個の並列処理が存在するとプログラムに伝えている。
	// 並列処理する個数が間違っているとエラーが起きる。
	wg.Add(1)
	// 並列処理
	// メインの処理とは別に並列に処理を実行している。
	// ※メインの処理が終了したら、並列処理を強制的に終了してしまう。
	// 変数wgのアドレスを渡す。
	go goroutine("world", &wg)
	// メイン処理
	normal("hello")
	// wg.Doneが呼ばれるまで処理を実行する。
	// 今回の処理ではgoroutineの動作が終了するまで待機すること。
	wg.Wait()
}