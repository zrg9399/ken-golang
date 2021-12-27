package main

import "fmt"

func main() {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done channel is triggerroed,exit chill go rountine")
				return
			}
		}

	}()
	close(done)

}
