package main

import (
	"fmt"
	"playground/boring"
)

func main() {

	// generator pattern used in Boring function
	result := boring.Boring("Joe")
	result1 := boring.Boring("Jane")

	for i := 0; i < 5; i++ {
		fmt.Println(<-result)
		fmt.Println(<-result1)
	}
	fmt.Println("You guys are talking too much")
}
