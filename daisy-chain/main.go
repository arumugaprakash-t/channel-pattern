package main

import "fmt"

func main() {

	max := 1000
	leftMost := make(chan int)
	left := leftMost
	var right chan int

	for i := 0; i < max; i++ {
		// create a new right channel pass it to the go routine then change it as left channel.
		// left <-right(left) <-right(left) <-right
		right = make(chan int)
		go chain(left, right)
		left = right
	}
	//send a value to the right channel, which will be propagated throughout the chain and result can be get from the left most channel
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftMost)
}

func chain(left, right chan int) {
	left <- 1 + <-right
}
