package main

import (
	"awesomeProject/mylib/lessons/s7_package"
	"awesomeProject/mylib/lessons/s7_package/under"
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	//goroutine.Goroutine09()
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s7_package.Average(s))
	s7_package.Say()
	under.Hello()

	person := under.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
	person.SetHeight(180)
	fmt.Println(person)
	fmt.Println(under.Public)

	/*
		https://github.com/markcheno/go-talib
		Exampleに従ってコピペ、試用
	*/
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
