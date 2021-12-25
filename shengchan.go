package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()
	//key:为自动生成的值，不需要手动输入
	ctx := context.WithValue(baseCtx, "a", "b")

	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)
	timeoutCtx, Cancel := context.WithTimeout(baseCtx, time.Second)
	defer Cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process imterrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)
	time.Sleep(1 * time.Second)
	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")

	}
}
