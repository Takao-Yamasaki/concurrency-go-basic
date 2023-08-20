package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// チャネルの読み込み
		fmt.Println(<-ch1)
	}()
	// チャネルへの書き込み
	ch1 <- 10
	// チャネルを閉じる
	close(ch1)
	// チャネルが開いているか閉じているか確認
	v, ok := <-ch1
	fmt.Printf("%v %v\n", v, ok)
	wg.Wait()
}
