package main

import (
	"fmt"
	"playground/boring"
)

func main() {

	a := boring.Boring("Joe")
	b := boring.Boring("Jane")

	result := fanIn(a, b)
	for i := 0; i < 20; i++ {
		fmt.Println(<-result)
	}

}

func fanIn(a, b <-chan string) <-chan string {
	result := make(chan string)
	go func() {
		for {
			result <- <-a
		}

	}()
	go func() {
		for {
			result <- <-b
		}
	}()
	return result
}
