package s8

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	defer close(ch)
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func Main02() {
	ch := make(chan string)
	//goroutineが時間がかかって終わらないような場合を想定してcontextを使右ことでタイムアウトを発生させられる。
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	//とりあえずタイムアウトも何もさせないならTODOを渡すということもできる。
	//ctx := context.TODO()
	defer cancel()
	go longProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("##################")
}
