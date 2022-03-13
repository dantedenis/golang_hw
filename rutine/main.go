package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	timeout := time.After(time.Millisecond)
	ch := make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < 300; i++ {
		wg.Add(1)
		go test(&wg, ch, i)
	}

	go func(ch chan int) {
		for {
			select {
			case <-timeout:
				fmt.Println("timeout")
				wg.Done()
				return
			default:
				fmt.Printf("%d", <-ch)
			}
		}
	}(ch)
	wg.Wait()
}

func test(wg *sync.WaitGroup, ch chan int, n int) {
	defer wg.Done()
	ch <- n
}
