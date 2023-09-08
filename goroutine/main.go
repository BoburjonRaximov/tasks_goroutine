package main

import (
	"fmt"
	"strings"
)

func main() {
	ch := make(chan string)

	go t2a(ch)
	//go t2b(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func t2a(ch chan<- string) {
	gaplar := []string{"Asssalomu aleykum", "ishlar yaxshimi ?", "uydagilar tinchmi ?"}
	belgi := "?"
	for i := 0; i < 3; i++ {
		result := strings.Contains(gaplar[i], belgi)
		if result == true {
			ch <- gaplar[i]
			break
		}
	}
	close(ch)
}

func t2b(ch chan<- string) {
	gaplar := []string{"Asssalomu aleykum", "ishlar yaxshimi ?", "uydagilar tinchmi ?"}
	belgi := "?"

	for i := 0; i < 3; i++ {
		result := strings.Contains(gaplar[i], belgi)
		if result == true {
			ch <- gaplar[i]
		}
	}
	close(ch)
}
