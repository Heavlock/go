package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	//ch <- 1
	//ch <- 2
	//ch <- 3
	go func() {
		//ch <- 4
		//close(ch)

		select {
		case val := <-ch:
			fmt.Println(555, val)
		default:
			return
		}

	}()
	time.Sleep(10 * time.Second)
	ch <- 5
	ch <- 5
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}

}
