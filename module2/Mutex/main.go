package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go rLock()
	go wLock()
	go Lock()
	time.Sleep(5 * time.Second)

}
func rLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rLock:", i)

	}

}
func wLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("wLock:", i)

	}

}

func Lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Lock:", i)

	}

}
