package main

import (
	"fmt"
	"time"
)

// 消费 chan 返回2个参数和返回1个，在close的情况下是不同的
func main() {
	loadChan := make(chan struct{})

	go func() {
		for {
			if _, ok := <-loadChan; ok {
				fmt.Println("ok")
			} else {
				fmt.Println("no ok")
			}
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 2)

	close(loadChan)

	time.Sleep(5 * time.Second)

}
