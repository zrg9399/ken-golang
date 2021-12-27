package main

import (
	"fmt"
)

func main() {
	//args := os.Args
	//fmt.Println("os args is:",os.Args)
	//name:=flag.String("name","world","specify the name you want to sya hi")
	//flag.Parse()
	//fmt.Println("name is:",name)
	//if len(args)!=0{
	//	fmt.Println("Do not accept any argument")
	//	os.Exit(1)
	//}
	fmt.Println("Hello World")

	fmt.Println(passValue(3, 4, 5))
}
func passValue(a int, b ...int) (x, y int) {
	x = a
	y = b
	return x, y

}
