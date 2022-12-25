package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg  sync.WaitGroup
	ch1 = make(chan int, 1)
	ch2 = make(chan int, 1)
	ch3 = make(chan int, 1)
)

func main() {
	ch1 <- 66

	wg.Add(3)
	go func() {
		<-ch1
		Work("goroutine1")
		ch2 <- 66
	}()
	go func() {
		<-ch2
		Work("goroutine2")
		ch3 <- 66
	}()
	go func() {
		<-ch3
		Work("goroutine3")
	}()
	wg.Wait()
	fmt.Println("successful")
}

func Work(workName string) {
	time.Sleep(time.Second) // 模拟逻辑处理
	fmt.Println(workName)
	wg.Done()
}
