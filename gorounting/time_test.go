package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	i := 0
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*4)
	for range time.Tick(time.Millisecond * 500) {
		select {
		case <-ctx.Done():
			cancelFunc()
			fmt.Println("退出")
			return
		default:
			fmt.Println("进入等待")
			i++
			if i == 3 {
				fmt.Println("退出")
				// 这里其实不会break出去
				break
			}
		}
	}
	fmt.Println("完成")
}

func TestTimeoutReturn(t *testing.T) {
	i := 0
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*6)
	for range time.Tick(time.Millisecond * 500) {
		select {
		case <-ctx.Done():
			cancelFunc()
			fmt.Println("退出")
			return
		default:
			i++
			fmt.Println("进入等待", i)
			time.Sleep(time.Second * 2)
			// 还会执行
			i++
		}
	}
	fmt.Println("完成")
}

func TestWithTimeout(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	go func() {
		for range time.Tick(time.Millisecond * 500) {
			t.Log("go func")
		}
		cancelFunc()
	}()
	<-ctx.Done()
	t.Log("结束")
}
