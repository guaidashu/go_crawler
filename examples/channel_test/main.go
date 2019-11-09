/**
  create by yy on 2019/11/9
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int, 1)

	block := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			queue <- i
			time.Sleep(time.Second * 10)
		}

		close(block)
	}()

	go func() {
		for v := range queue {
			fmt.Println(v)
		}
	}()

	<-block

}
