package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("test1")
	defer fmt.Println("test2")
	defer fmt.Println("test3")
	loopFunc()
	time.Sleep(time.Second)

}
func loopFunc() {
	lock := sync.Mutex{}
	for i := 1; i <= 3; i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopFunc:", i)
		}(i)
	}
}
