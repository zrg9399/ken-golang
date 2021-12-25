package main

import (
	"fmt"
	//"math/rand"
	//"time"
	"errors"
)

func main() {
	//异常处理
	var errorNotFond error = errors.New("NotFound")
	fmt.Println(errorNotFond)
	//err := fmt.Errorf( "this is error")
	//fmt.Println(err)
	// 多协程
	//go fmt.Println(11)
	//go fmt.Println(22)
	//go fmt.Println(33)
	//for i := 0;i<10;i++ {
	//	go fmt.Println(i)
	//}
	//time.Sleep(time.Second)
	//数组替换
	//mySlice := []string{"I", "am", "stupid", "and", "weak"}
	//
	//for index, _ := range mySlice {
	//	if mySlice[index] == "stupid" {
	//		mySlice[index] = "smart"
	//	} else if mySlice[index] == "weak"  {
	//		mySlice[index] = "strong"
	//
	//	}
	//	fmt.Println(mySlice[index])
	//}
	//**********通道***********************
	//ch := make(chan int)
	//go func() {
	//	fmt.Println("this is chird thread")
	//	ch <- 1
	//}()
	//<- ch
	//	********************遍历通道缓冲区******************
	//ch := make(chan int,2)
	//ch1 :=make(chan int,2)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		//纳秒级随机种子,生成随机数
	//		rand.Seed(time.Now().UnixNano())
	//		n := rand.Intn(1000)
	//		fmt.Println("putting", n)
	//		ch <- n
	//	}
	//	//告诉接受者没有新数据
	//	close(ch)
	//}()
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		//纳秒级随机种子,生成随机数
	//		rand.Seed(time.Now().UnixNano())
	//		n := rand.Intn(1000)
	//		fmt.Println("putting", n)
	//		ch1 <- n
	//	}
	//	//告诉接受者没有新数据
	//	close(ch)
	//}()
	//fmt.Println("Hello form main")
	//for v :=range ch{
	//	fmt.Println("receiving",v)
	//}

}
