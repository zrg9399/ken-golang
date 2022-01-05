package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	productCount = 5
	mutex        sync.Mutex
	cond         = sync.NewCond(&mutex)
)

type producer struct {
}

func (r *producer) produce() {
	for {
		cond.L.Lock()
		if productCount < 10 {
			productCount++
			fmt.Printf("生产者生产了一个产品，当前存量%d\n", productCount)
		} else {
			fmt.Printf("仓库满了，当前容量%d\n", productCount)
			cond.Wait()
		}
		cond.L.Unlock()
		cond.Broadcast()
	}
}
func (r *consumer) consumer() {
	for {
		cond.L.Lock()
		if productCount > 0 {
			productCount--
			fmt.Printf("消费者消费了一个产品，当前存量%d\n", productCount)
		} else {
			fmt.Printf("仓库空了，当前存量%d\n", productCount)
			cond.Wait()
		}
		cond.L.Unlock()
		cond.Broadcast()
	}
}

type consumer struct {
}

func main() {
	var p = producer{}
	var c = consumer{}
	go p.produce()
	go c.consumer()
	time.Sleep(time.Microsecond * 10000)
}
