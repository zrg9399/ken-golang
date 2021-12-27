package main

////
////import (
////	"fmt"
////)
////
////func main() {
////	c := make(chan int )
////	go Producer(c)
////	Consumer(c)
////}
////
////func Producer(c chan int) {
////	for i := 0; i < 10; i++ {
////		//fmt.Println("Produce:", i)
////		c <- i
////	}
////	close(c)
////}
////
////func Consumer(c chan int) {
////	for num := range c {
////		fmt.Println("Consume:", num)
////	}
////}
//import ( (
//	"fmt"mt"
//	"sync"ath/rand"
//)ime"
//var (
//	productCount = 5
//	mutex sync.MutexuntChan chan int
//	cond = sync.NewCond(&mutex)
//)roducer struct {
//type producer struct {
//
//}r *producer) produce() {
//func (r *producer) produce() {r i := 0; i < 10; i++ {
//	for {	select {
//		cond.L.Lock()	case count := <-countChan:
//		if productCount < 10 {		if count < 10 {
//			productCount++			count++
//			fmt.Printf("生产者生产了一个产品，当前存量%d\n", productCount)			fmt.Printf("生产者生产了一个产品，当前存量%d\n", count)
//		} else {		} else {
//			fmt.Printf("仓库满了，当前容量%d\n", productCount)			fmt.Printf("仓库满了，当前容量%d\n", count)
//			cond.Wait()		}
//		}		countChan <- count
//		cond.L.Unlock()		var t = rand.Intn(100)
//		cond.Broadcast()		time.Sleep(time.Duration(t))
//	}	default:
//}		fmt.Println("初始化产品池")
//func (r *consumer) consumer() {		countChan <- 0
//	for {	}
//		cond.L.Lock()
//		if productCount > 0 {
//			productCount--
//			fmt.Printf("消费者消费了一个产品，当前存量%d\n", productCount)r *consumer) consumer() {
//		} else {r i := 0; i < 10; i++ {
//			fmt.Printf("仓库空了，当前存量%d\n", productCount)	select {
//			cond.Wait()	case count := <-countChan:
//		}		if count > 0 {
//		cond.L.Unlock()			count--
//		cond.Broadcast()			fmt.Printf("消费者消费了一个产品，当前存量%d\n", count)
//	}		} else {
//}			fmt.Printf("仓库空了，当前存量%d\n", count)
//type consumer struct {		}
//		countChan <- count
//}		var t = rand.Intn(100)
//func main() {		time.Sleep(time.Duration(t))
//	var p = producer{}	default:
//	var c = consumer{}		fmt.Println("初始化产品池")
//	go p.produce()		countChan <- 0
//	go c.consumer()	}
//}
//
//
//
//onsumer struct {
//
//
//ain() {
//untChan = make(chan int, 1)
//r p = producer{}
//r c = consumer{}
// p.produce()
// c.consumer()
//me.Sleep(10000)
//
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
