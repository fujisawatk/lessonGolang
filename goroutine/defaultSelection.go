package main

import (
		"fmt"
		"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	// Label
	OuterLoop:
		for {
			select {
			case <-tick:
				fmt.Println("tick.")
			case <-boom:
				fmt.Println("Boom!")
				// returnの場合：main関数自体を抜けてしまう。
				// breakの場合：selectのboomが中断されるので、
				// for文のループは止まらない。（tickとdefaultでループし続ける）

				// Labeled Break
				// →ラベル名を指定してbreakすることで、ラベル直下のfor文を中断できる。
				// （終了するのはfor文のみで、main関数自体は終了しない）
				break OuterLoop
			default:
				fmt.Println("    .")
				time.Sleep(50 * time.Millisecond)
			}
		}
		// Labeled Breakの次から処理を続行する。
		fmt.Println("################")
}