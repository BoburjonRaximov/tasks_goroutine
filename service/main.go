package main

import "fmt"

func main() {
	message("hello", "Adam")
	message("hello", "nodir")

}

func message(text, receiver string) {
	blac_list := []string{"Adam", "John", "Tom"}
	var m string
	for i := 0; i < 3; i++ {
		if receiver == blac_list[i] {
			m = "Message send"
			fmt.Println(m)
		}
	}
	if m != "Message send" {
		fmt.Println("Message ignored")
	}

}
