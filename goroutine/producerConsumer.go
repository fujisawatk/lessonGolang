package main

import (
			"sync"
			"fmt"
			"time"
)

func producer(ch chan int, i int){
	// Something
	// 処理をしてchannelに送信
	// chに10個値が格納される。
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup){
	// channelの値の数分for$range文が動作
	for i := range ch{
		// inner function
		func (){
			// inner functionの処理が全部終わったら実行
			defer wg.Done()
			fmt.Println("process", i * 1000)
		}()
	}
}

func main(){
	// syncパッケージWaitGroupメソッドを宣言
	var wg sync.WaitGroup
	// channel宣言
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		// 1個の並列処理を伝達
		// 10回繰り返し文なので、10個の並列処理が存在すると伝達している。
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer
	// producerメソッドでchannelに格納した値、WaitGroupを引数で持っていく。
	go consumer(ch, &wg)
	// consumerメソッドの処理が終わるまで待機
	wg.Wait()
	// channelのデータ操作が全部終了したら閉じる宣言
	close(ch)
}