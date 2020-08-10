package main

import (
	"fmt"
	"time"
)

func main() {
	//当多个需要从多个chan中读取或写入时，会先轮询一遍所有的case，
	//然后在所有处于就绪（可读/可写）的chan中随机挑选一个进行读取或写入操作，
	//并执行其语句块。如果所有case都未就绪，则执行default语句，
	//如未提供default语句，则当前协程被阻
	select {
	case resp := <-AsyncCall(50):
		fmt.Println(resp)
	case resp := <-AsyncCall(200):
		fmt.Println(resp)
	case resp := <-AsyncCall2(3000):
		fmt.Println(resp)
	}
	//这段代码运行的结果会是200和50两种结果随机出现。
}

func AsyncCall(t int) <-chan int {
	c := make(chan int, 1)
	go func() {
		time.Sleep(time.Microsecond * time.Duration(t))
		c <- t
	}()
	return c
}

func AsyncCall2(t int) <-chan int {
	c := make(chan int, 1)
	go func() {
		time.Sleep(time.Microsecond * time.Duration(t))
		c <- t
	}()
	// gc or some other reason cost some time
	time.Sleep(200 * time.Microsecond)
	return c
}
