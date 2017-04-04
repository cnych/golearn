package main

import (
	"bytes"
	"fmt"
	"sync"
)

// 在一些复杂的程序中，通常通过不同线程执行不同应用来实现程序的并发。
// 当不同线程要使用同一个变量时，经常会出现一个问题：无法预知变量被不同线程修改的顺序！
// 这通常被称为资源竞争，指不同线程对同一变量使用的竞争。
// 经典的嗯做法是一次只能让一个线程对共享变量进行操作。当变量被一个线程改变时，我们为它上锁，
// 直到这个线程执行完成并解锁后，其他线程才能访问它。
// 我们之前章节学习的map类型是不存在锁的机制来实现这种效果的，所以map类型是非线程安全的。
// 当并行访问一个共享的map类型的数据，map数据将会出错。
// 在golang中通过sync包中的Mutex来实现锁。
// sync.Mutex 是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。
// 假如info是一个需要上锁的放在共享内存中的变量,可以通过包含Mutex来实现：
type Info struct {
	mu  sync.Mutex
	Str string
	// ...other fields.
}

// 如果一个函数想要改变这个变量
func Update(info *Info) {
	info.mu.Lock()
	info.Str = "updating~"
	info.mu.Unlock()
}

// 通过Mutex来实现一个可以上锁的共享缓冲器
type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

// sync 包中还有一个RWMutex锁：他能通过RLock()来允许同一时间多个线程对变量进行读操作
// 但是只能有一个线程进行写操作。如果使用Lock()将和普通的Mutex作用相同。包中还有一个方便的Once类型变量的方法
// once.Do(call),这个方法确保被调用函数只能被调用一次。

// 相对简单的情况下，通过使用sync包可以解决同一时间只能一个线程访问变量或map类型数据的问题。
// 如果这种方式导致程序明显变慢或者引起其他问题，我们要重新思考来通过goroutine和channel来解决了
func main() {
	fmt.Println("sync.Mutex用法~")
	info := Info{Str: "Hello World"}
	fmt.Printf("info=%v\n", info)
	Update(&info)
	fmt.Printf("info=%v\n", info)
}
