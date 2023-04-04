package s6

import "fmt"

func goroutine031(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		// forループ中都度結果を呼び出し元に返していきたい時、ここにchanを使う
		c <- sum
	}
	close(c)
}

func Main03() {
	s := []int{1, 3, 5, 7, 9}
	// chanの容量を指定せず無制限にもできるがこれで最小にできる
	c := make(chan int, len(s))
	go goroutine031(s, c)
	// cを受け取る度に処理が走る
	for i := range c {
		fmt.Println(i)
	}
}
