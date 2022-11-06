// lv2
package main

import "fmt"

func main() {
	var (
		in  = make(chan struct{}, 1)
		out = make(chan struct{}, 2)
	)

	go MyPrint(in, out, 2)
	go MyPrint(in, out, 1)

	for i := 0; i < 2; i++ {
		<-out
	}
}

func MyPrint(in, out chan struct{}, i int) {
	for ; i <= 100; i += 2 {
		in <- struct{}{}
		fmt.Println(i)
		<-in
	}
	out <- struct{}{}
}
