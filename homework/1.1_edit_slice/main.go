package main

import "fmt"

func main() {
	//修改数组值
	mySlice := []string{"I", "am", "stupid", "and", "weak"}
	for index, _ := range mySlice {
		if mySlice[index] == "stupid" {
			mySlice[index] = "test"
		} else if mySlice[index] == "weak" {
			mySlice[index] = "strong"
		}

		fmt.Println(mySlice[index])

	}

}
