package s6

import "fmt"

func Main02() {
	fmt.Println("goroutine02 start")

	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	// ch を 2 で作ったので3回目はエラーになる。
	//ch <- 300
	//fmt.Println(len(ch))

	// x に ch の値を一つ取り出すと空きが一つできるのでもう一度chを使える。この後４回目は使えない。
	x := <-ch
	fmt.Println(x)

	fmt.Println(len(ch))

	ch <- 300
	fmt.Println(len(ch))

	// channelをクローズすると使い直せる。この前にforを回しても3つめから取り出そうとしてエラーになる。
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}
}
