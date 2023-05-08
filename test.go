package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(5 * 3)
	for i := 0; i < 5; i++ {
		go gorutine(ch)
	}
	//for i := 0; i < 16; i++ {
	go reader(ch, wg)
	//}
	//wg.Wait()
	fmt.Println("end")
}
func gorutine(ch chan<- string) {
	ch <- "str"
	ch <- "str2"
	ch <- "str"
}
func gorutine2(ch <-chan string, wg *sync.WaitGroup) {
	//ch <- "gorutine2"

}
func reader(ch <-chan string, group *sync.WaitGroup) {
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
			group.Done()
		}
	}

}
