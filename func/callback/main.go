package main

import "fmt"

func main() {
	DoOperation(2, increase)
	DoOperation(2, decrease)

}
func increase(a, b int) {
	fmt.Println("increase result is:", a+b)

}
func DoOperation(y int, f func(int, int)) {
	f(y, 2)

}
func decrease(a, b int) {
	fmt.Println("increase result is:", a-b)

}
