package s6

import "fmt"

func goroutine011(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func goroutine012(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Main01() {
	fmt.Println("goroutine01 start")

	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine011(s, c)
	go goroutine012(s, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
