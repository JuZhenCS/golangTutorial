package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup //同步等待

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1) //在主进程加任务数
		go child(i)
	}
	waitGroup.Wait()
}

func child(number int) {
	//waitGroup.Add(1)  //不能在子进程加任务数，否则还没加上，主任务就结束了
	defer waitGroup.Done() //任务完成
	fmt.Println("I am child ", number)
}
