package s6

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, i int) {
	// 何かの処理を作成
	ch <- i * 2
}

// producerが並行処理でchに渡した値をgoroutineであるconsumerが受け取ってrange処理する。
// goroutine03でmain処理でしていたことを別のgoroutineで実行している形。
func consumer041(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			// インナーfuncが実行後必ずdeferが実行されるため、処理が何かの原因で飛んだとしてもdeferは実行される。
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
	fmt.Println("#####################")
}

func Goroutine04() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer
	go consumer041(ch, &wg)
	// producer一回の処理にconsumerのfor文1周を対応させる。
	wg.Wait()
	// ここでcloseしないと次のchを待ち続けてしまい、以降の処理（consumerのfor文よりあと）が実行されない。
	close(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
