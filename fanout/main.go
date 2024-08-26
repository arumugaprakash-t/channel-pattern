package main

import (
	"fmt"
	"playground/boring"
	"time"
)

func main() {

	a := boring.Boring("Joe")
	b := boring.Boring("Jane")

	result := fanIn(a, b)

	fanOutResult := fanOut(result, 5)
	timeOut := time.After(10 * time.Second)
	for {
		select {
		case value := <-fanOutResult:
			fmt.Println(value)
		case <-timeOut:
			fmt.Println("Worked enough today")
			return
		}
	}
}

func fanOut(input <-chan string, workers int) <-chan string {
	output := make(chan string)
	for i := 0; i < workers; i++ {
		go worker(input, output)
	}
	return output
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

func worker(input <-chan string, output chan string) {
	for work := range input {
		fmt.Printf("Got %s\n", work)
		time.Sleep(10 * time.Millisecond)
		output <- fmt.Sprintf("Completed %s\n", work)
	}

}
