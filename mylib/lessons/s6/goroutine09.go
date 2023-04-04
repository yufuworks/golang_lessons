package s6

import "fmt"

func goroutine091(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	//closeしないと値の受け取り待機を続けてしまいデッドロックになり渡した結果を出力できない。
	close(c)
}

func Main09() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	//文字列を返すchanを作成する必要がある。
	c := make(chan string)
	go goroutine091(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
