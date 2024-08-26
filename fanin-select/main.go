package main

import (
	"fmt"
	"playground/boring"
	"time"
)

func main() {

	a := boring.Boring("Joe")
	b := boring.Boring("Jane")

	result := fanInSelect(a, b)
	timeOut := time.After(5 * time.Second)
	for {
		select {
		case <-timeOut:
			fmt.Println("You guys are talking too much..")
			return
		case value := <-result:
			fmt.Println(value)

		}
	}

}

func fanInSelect(a, b <-chan string) <-chan string {
	result := make(chan string)
	go func() {
		for {
			select {
			case value := <-a:
				result <- value
			case value := <-b:
				result <- value
			}
		}
	}()
	return result
}
