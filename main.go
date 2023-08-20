package main

import (
	"fmt"
	"runtime"
	// "sync"
	// "time"
)

func main() {
	// 	ch := make(chan int)
	// 	var wg sync.WaitGroup
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		ch <- 10
	// 		time.Sleep(500 * time.Millisecond)
	// 	}()
	// 	fmt.Println(<-ch)
	// 	wg.Wait()

	ch1 := make(chan int)
	go func() {
		// 読み込み操作を行なっているが、書き込みがないので、ずっと待機している
		fmt.Println(<-ch1)
	}()
	// 書き込みを追加してみる
	ch1 <- 10
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())

	// バッファ付きchannel
	ch2 := make(chan int, 1)
	// 書き込み
	ch2 <- 2
	ch2 <- 3
	// 読み込み
	fmt.Println(<-ch2)
}
