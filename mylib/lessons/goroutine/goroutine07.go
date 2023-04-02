package goroutine

import (
	"fmt"
	"time"
)

func Goroutine07() {
	//Tick, Afterはtimeが持つchannel
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	//ループに名前をつけてbreakで何を抜けるかを指定できる...golandでインデントが自動で戻されてしまう、なぜ？？
Outerloop:
	for {
		select {
		//tickには時間情報が渡されてくるので取り出し可能。Afterも同様
		case <-tick:
			//case t := <-tick:
			fmt.Println("tick.")
			//fmt.Println("tick.", t)
		case <-boom:
			//case a := <-boom:
			fmt.Println("BOOM!")
			//fmt.Println("BOOM!", a)
			//breakだと止まらない
			break Outerloop
			//return
		//defaultの処理はchannelとは無関係
		default:
			fmt.Println("      .")
			time.Sleep(50 * time.Millisecond)
		}
		//ここにbreakをおくと最初のcaseを処理して終わってしまう
	}
	fmt.Println("##############################")
}
