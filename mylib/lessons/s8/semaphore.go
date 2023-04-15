package s8

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// goroutineの同時実行数を制限できる、並列処理の制御ができる。
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess080501(ctx context.Context) {
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)
	fmt.Println("01 Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("01 Done")
}

func longProcess080502() {
	// 指定されたsemaphore数のgoroutineが走っていた場合falseとなってロック状態になるため、他の処理がキャンセルされる。
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("02 Could not get lock")
		return
	}
	defer s.Release(1)
	fmt.Println("02 Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("02 Done")
}

func Main0501() {
	ctx := context.TODO()
	go longProcess080501(ctx)
	go longProcess080501(ctx)
	go longProcess080501(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("01 main Done")
}

func Main0502() {
	go longProcess080502()
	go longProcess080502()
	go longProcess080502()
	time.Sleep(5 * time.Second)
	fmt.Println("02 main Done")
}
