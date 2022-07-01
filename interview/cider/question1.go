package cider

import (
	"fmt"
	"runtime"
	"sync"
)

//写一个死锁
//情况一：无缓存能力的通道，自己写自己读
func Deadlock1() {
	ch := make(chan int)
	ch <- 666
	x := <-ch
	fmt.Println(x)
}

//情况二：协程来晚了
func Deadlock2() {
	ch := make(chan int)
	ch <- 666
	go func() {
		<-ch
	}()
}

//情况三 管道读写时，互相要求对方先读/写
func Deadlock3() {
	chHusband := make(chan int)
	chWife := make(chan int)
	go func() {
		if _, ok := <-chHusband; ok {
			chWife <- 888
		}
		// select {
		// case <-chHusband:
		// 	chWife <- 888
		// }
	}()
	// select {
	// case <-chWife:
	// 	chHusband <- 666
	// }
	if _, ok := <-chWife; ok {
		chHusband <- 888
	}
}

//情况四：读写锁相互阻塞，形成隐形死锁
func Deadlock4() {
	var rmw09 sync.RWMutex
	ch := make(chan int)
	go func() {
		rmw09.Lock()
		ch <- 999
		rmw09.Unlock()
	}()
	go func() {
		rmw09.RLock()
		x := <-ch
		fmt.Println("读到", x)
		rmw09.RUnlock()
	}()
	for {
		runtime.GC()
	}
}
