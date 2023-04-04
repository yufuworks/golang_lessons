package s6

import (
	"fmt"
	"time"
)

func goroutine61(ch chan string) {
	//無限ループで1秒ごとにchに文字列を送る
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func goroutine62(ch chan string) {
	//無限ループで1秒ごとにchに文字列を送る
	for {
		ch <- "packet from 1"
		time.Sleep(2 * time.Second)
	}
}

func Main06() {
	c1 := make(chan string)
	c2 := make(chan string)
	go goroutine61(c1)
	go goroutine62(c2)

	for {
		//c1, c2 という別々のchan処理（goroutineの結果）を同時に扱える
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
