package main

import "fmt"

func fibona() func(int) int {
	i, j := 0, 1
	return func(k int) int {
		if k > 0 {
			i, j = j, j+i
		}
		return i
	}

}
func main() {
	fibonaa := fibona()
	for i := 0; i < 10; i++ {
		fmt.Println(fibonaa(i))
		//fmt.Println("i=", i, "num=", fibonaa(i))
	}

}
