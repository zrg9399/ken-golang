package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//*******************遍历缓冲区数据****************
	ch := make(chan int, 10)
	ch1 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			//纳秒级随机种子,生成随机数
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(1000)
			fmt.Println("putting", n)
			ch <- n
		}
		//告诉接受者没有新数据
		close(ch)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			//纳秒级随机种子,生成随机数
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(1000)
			fmt.Println("putting", n)
			ch1 <- n
		}
		//告诉接受者没有新数据
		close(ch)
	}()
	fmt.Println("Hello form main")
	for v := range ch {
		fmt.Println("receiving", v)
	}
	//*****************生产者&消费者问题***************************
	//messages := make(chan int, 10)
	//done := make(chan bool)
	//defer close(messages)
	////consumer
	//go func() {
	//	ticker := time.NewTimer(1 * time.Second)
	//	for _ = range ticker.C {
	//		select {
	//		case <-done:
	//			fmt.Println("child process interrupt...")
	//			return
	//		default:
	//			fmt.Printf("send message: %d\n", <-messages)
	//
	//		}
	//	}
	//	//producer
	//	for i := 0; i < 10; i++ {
	//		messages <- i
	//	}
	//	time.Sleep(5 * time.Second)
	//	close(done)
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("main process exit!")
	//}()

}
