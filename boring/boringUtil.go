package boring

import (
	"fmt"
	"time"
)

func Boring(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s says := %d", name, i)
			time.Sleep(1 * time.Second)
		}
	}()
	return c
}
