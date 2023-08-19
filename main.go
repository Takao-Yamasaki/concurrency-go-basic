package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	// errGroupの作成
	eg := new(errgroup.Group)
	s := []string{"task1", "fake1", "task2", "fake2"}
	for _, v := range s {
		task := v
		// Goを使用することでエラーを受け取ることができる
		eg.Go(func() error {
			return doTask(task)
		})
	}
	// errGroupのWaitを使用することで、ゴルーチンが終了するまで待機
	// 最後まで実行して、最初に検知されたエラーを返す
	if err := eg.Wait(); err != nil {
		fmt.Printf("error :%v\n", err)
	}
	fmt.Println("finish")
}

func doTask(task string) error {
	if task == "fake1" || task == "fake2" {
		return fmt.Errorf("%v failed", task)
	}
	fmt.Printf("task %v completed\n", task)
	return nil
}
