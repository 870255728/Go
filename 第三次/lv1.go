package main

import (
	"fmt"
	"sync"
)

// lv1
var (
	x  int64
	wg sync.WaitGroup
)

func main() {
	var ch = make(chan struct{}, 1)
	wg.Add(2)
	go add(ch)
	go add(ch)
	wg.Wait()
	fmt.Println(x)
}

func add(ch chan struct{}) {
	for i := 0; i < 50000; i++ {
		ch <- struct{}{}
		x = x + 1
		<-ch
	}
	wg.Done()
}
