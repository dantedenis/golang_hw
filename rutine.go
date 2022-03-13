package main

import (
    "fmt"
    "sync"
    )

func main() {
    ch := make(chan int)
    wg := sync.WaitGroup{}
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go test(&wg, ch, i)
    }
    for i := range ch {
        fmt.Println(i)
    }
    //for i := 0; i < 5; i++ {  ----->>> not deadlock
     //   fmt.Println(<-ch)
    //}
    wg.Wait()
}

func test(wg *sync.WaitGroup, ch chan int, n int){
    defer wg.Done()
    ch <- n
}
