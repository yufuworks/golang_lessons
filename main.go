package main

import (
	"awesomeProject/mylib/lessons/s7_package"
	"awesomeProject/mylib/lessons/s7_package/under"
	"fmt"
)

func main() {
	//goroutine.Goroutine09()
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s7_package.Average(s))
	s7_package.Say()
	under.Hello()
}
