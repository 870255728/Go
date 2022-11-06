package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		//这里要么用WaitGroup要么用管道来保证每一个数字都被打印完才结束main groutine
		wg sync.WaitGroup
		//需要初始化一个带有一个缓存的管道，以防在if判断语句的时候整个程序阻塞直接报错
		over = make(chan bool, 1)
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		if i == 9 {
			over <- true
		}
	}
	wg.Wait()
	<-over
	fmt.Println("over!!!")
}
