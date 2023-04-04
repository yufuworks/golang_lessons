package s6

import "fmt"

func producer051(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func multi2(first chan int, second chan int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

// 明示的に chan の出入りを引数の中で表記できる。扱いがわかりやすくなる。
// func multi4(second <-chan int, third chan<- int) {
func multi4(second chan int, third chan int) {
	defer close(third)
	for i := range second {
		third <- i * 4
	}
}

func Main05() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	// 3つのchanを使ってそれぞれのgoroutineの処理を順次渡して処理している。
	go producer051(first)
	go multi2(first, second)
	go multi4(second, third)
	for result := range third {
		fmt.Println(result)
	}

}
