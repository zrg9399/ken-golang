package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//map转化为json
func main() {
	map_json()
	array_json()
	string_interface_json()
	struct_json()
	stuct_a_json()
	josn_map()
}
func map_json() {
	user := make(map[string]string)
	user["username"] = "kongyixueyuan"
	user["address"] = "北京"
	jsonStr, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", jsonStr)
}

//array转换成json
func array_json() {

	personArr := [3]string{"小李", "王二狗", "老毕"}
	jsonStr, err := json.Marshal(personArr)
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("%s\n", jsonStr)
}
func string_interface_json() {
	personArr := [3]interface{}{"张飞", "关羽", "刘备"}
	jsonStr, err := json.Marshal(personArr)
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("%s\n", jsonStr)

}

//结构体转换json
func struct_json() {

	type Person struct {
		//结构体首字母大写
		UserName string
		Age      int
	}
	str := Person{"王二狗", 30}
	jsonStr, err := json.Marshal(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonStr))

}

//************************嵌套结构体转换json***************
type Address struct {
	Province string
	City     string
	Number   int
}

type Person struct {
	UserName string `json:"aaa"`
	Age      int    `json:"-"`
	Addrs    Address
}

func stuct_a_json() {
	str := Person{"王二狗", 30, Address{"上海", "浦东新区", 1803}}
	jsonStr, err := json.Marshal(str)
	//
	jsonStr2, err := json.MarshalIndent(str, "+", "-")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonStr))
	fmt.Println(string(jsonStr2))

}

//json转换map
func josn_map() {
	str := "{\"address\":\"北京\",\"username\":\"kongyixueyuan\"}"
	myMap := make(map[string]string)
	json.Unmarshal([]byte(str), &myMap)
	fmt.Println(myMap)
}
