package main

import (
	"fmt"
	"os"
	"time"
)

var ch1 chan int = make(chan int)
var bufChan chan int = make(chan int, 1000)
var msgChan chan string = make(chan string)

func sum(a int, b int) {
	ch1 <- a + b
}

// write data to channel
func writer(max int) {
	for {
		for i := 0; i < max; i++ {
			bufChan <- i
			fmt.Fprintf(os.Stderr, "%v write: %d\n", os.Getpid(), i)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// read data fro m channel
func reader(name string) {
	for {
		r := <-bufChan
		fmt.Printf("%s read value: %d\n", name, r)
	}
	msgChan <- name
}

func testWriterAndReader(max int) {
	// 开启多个writer的goroutine，不断地向channel中写入数据
	go writer(max)
	go writer(max)

	// 开启多个reader的goroutine，不断的从channel中读取数据，并处理数据
	go reader("read1")
	go reader("read2")
	go reader("read3")

	// 获取三个reader的任务完成状态
	name1 := <-msgChan
	name2 := <-msgChan
	name3 := <-msgChan

	fmt.Println("%s,%s,%s: All is done!!", name1, name2, name3)
}

func main() {
	testWriterAndReader(100)
}
