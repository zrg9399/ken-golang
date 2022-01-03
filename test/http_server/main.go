package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	get()
	post()
	post_json()
	post_client()

}

func get() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
	req.Header.Add("name", "zhaofan")
	req.Header.Add("age", "3")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}

func post() {
	urlValues := url.Values{}
	urlValues.Add("name", "zhaofan")
	urlValues.Add("age", "22")
	resp, _ := http.PostForm("http://httpbin.org/post", urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func post_json() {
	client := &http.Client{}
	data := make(map[string]interface{})
	data["name"] = "zhaofan"
	data["age"] = "23"
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "http://httpbin.org/post_json", bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func post_client() {
	data := make(map[string]interface{})
	data["name"] = "zhaofan"
	data["age"] = "23"
	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post("http://httpbin.org/post_client", "application/json", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
