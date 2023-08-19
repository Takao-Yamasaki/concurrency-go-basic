package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// カウンタを1つ加算
	wg.Add(1)
	go func() {
		// カウンタを1つ減算
		defer wg.Done()
		fmt.Println("goroutine invoked")
	}()
	wg.Wait()
	// 起動しているゴルーチンの数を取得する
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("main func finish")
}
