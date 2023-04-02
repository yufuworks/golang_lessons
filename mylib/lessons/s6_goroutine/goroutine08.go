package s6_goroutine

import (
	"fmt"
	"sync"
	"time"
)

func Goroutine08_1() {
	/*
		このまま実行すると2つのgoroutineが同時に同じmapを書き換えようとしてバッティングしてエラーが発生することがある。
	*/
	c := make(map[string]int)
	go func() {
		for i := 0; i < 10; i++ {
			c["key"] += 1
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c["key"] += 1
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c["key"])
}

type counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *counter) Inc(key string) {
	/*
		この関数を実行するとき、まずプロセスにロックをかける。
		必要な操作（v : map の key : 引数で指定された文字列に対するkey の value をインクリメント）
		をした後最後にプロセスをアンロックし、他からの呼び出しを許可する。
	*/
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func Goroutine08_2() {
	/*
		mutexを使うことでアクセス制御できる構造体ができ、アクセスがバッティングしてエラーになることがなくなる。
	*/
	//c := make(map[string]int)
	c := counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			//c["key"] += 1
			c.Inc("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			//c["key"] += 1
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, ":", c.Value("key"))
}
