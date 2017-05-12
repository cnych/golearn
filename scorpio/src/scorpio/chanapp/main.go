package main

import "time"
import "fmt"

func main() {
	// 无缓冲的与有缓冲Channel有着很大的差别，那就是一个是同步的，一个是非同步的。
	// c1 := make(chan int)  // 无缓冲
	// c2 := make(chan int, 1)  // 有缓冲
	// c1 <- 1  // 不仅仅是向c1通道放1，而是一直要等有别的携程<-c1接收这个参数，那么c1<-1才会继续下去，要不然就一直阻塞着。
	// c2 <- 2 // 不会阻塞，因为缓冲大小是1，只有当放第二个值的时候，第一个还没被人拿走，这时候才会阻塞。
	// test0()
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	// test9()
	// test10()
	// test11()
	// test12()
	// test13()
	// test14()
	// test15()
	// test16()
	test17()
}

func test0() {
	done := make(chan bool)
	done1 := make(chan bool, 1)
	println(done, done1)
}

func test1() {
	// 编译错误 deadlock，阻塞main进程
	done := make(chan bool)
	done <- true // 这句是输入值，它会一直阻塞，等待读取
	<-done       // 这句是读取，但是在上面已经阻塞了，永远走不到这里
	println("完成")
}

func test2() {
	// 同样deadlock，仅有输入语句，没有读取语句的死锁
	done := make(chan bool)
	done <- true
	println("完成")
}

func test3() {
	// deadlock，仅有读取，没有输入语句的死锁
	done := make(chan bool)
	<-done // 读取输出，前面没有输入语句，done是empty的，所以一直等待输入
	println("完成")
}

func test4() {
	// 编译通过，协程阻死，不会影响main进程
	done := make(chan bool)
	go func() {
		<-done // 一直等待
	}()
	println("完成")
}

func test5() {
	done := make(chan bool)
	go func() {
		println("我肯能会输出哦") // 阻塞前的语句
		done <- true       // 这里阻塞死，但是上面的语句可能会出现
		println("我永远不会输出")
		<-done // 这句也不会走到，除非在别的协程里面读取，或者在main里面
	}()
	println("完成")
}

func test6() {
	/** 编译通过，在 test5 的基础上演示，延时 main 的跑完 */
	done := make(chan bool)
	go func() {
		println("我可能会输出哦")
		done <- true /** 这里阻塞死 */
		println("我永远不会输出")
		<-done /** 这句也不会走到 */
	}()
	time.Sleep(time.Second * 1) /** 加入延时 1 秒 */
	println("完成")
	/**
	 * 控制台输出：
	 *       我可能会输出哦
	 *       完成
	 */
	/**
	 * 结论：
	 *    如果在 go routine 中阻塞死，也可能不会把阻塞语句前的内容输出，
	 *    因为main已经跑完了，所以延时一会，等待 go routine
	 */
}

func test7() {
	done := make(chan bool)
	go func() {
		done <- true
		println("我永远不会输出，除非<-done执行")
	}()
	<-done // 这里接收，在输出完成之前，上面的语句将会走通
	println("完成")
}

func test8() {
	done := make(chan bool)
	go func() {
		done <- true
		println("我永远不会输出，除非<-done执行")
	}()
	println("完成")
	<-done // 这里接收，在输出完成之后
}

func test9() {
	// 没有缓冲的channel使用close后，不会阻塞
	done := make(chan bool)
	close(done)
	//done<-true // 关闭了，不能再往里面输入值
	<-done // 这句是读取，但是在上面已经关闭channel了，不会阻塞
	println("完成")
}

func test10() {
	// 没缓冲的channle，在协程中close后，不会阻塞
	done := make(chan bool)
	go func() {
		close(done)
	}()
	// done<-true // 关闭了，不能再往里面输入值了
	<-done // 这句是读取，但是在上面已经关闭channel了，不会阻塞
	println("完成")
}

func test11() {
	// 有缓冲的channel不会阻塞
	done := make(chan bool, 1)
	done <- true
	<-done
	println("完成")
}

func test12() {
	// 有缓冲channle会阻塞的例子
	done := make(chan bool, 1)
	// done<-true
	<-done // 在没有输入的情况下，读取，会阻塞
	println("完成")
}

func test13() {
	// 有缓冲channel会阻塞的例子
	done := make(chan bool, 1)
	done <- true
	done <- false //放第二个值的时候，第一个还没被人拿走，这时候才会阻塞
	println("完成")
}

func test14() {
	// 有缓冲channel不会阻塞的例子
	done := make(chan bool, 1)
	done <- true // 不会阻塞，等待读取
	println("完成")
}

func test15() {
	// 有缓冲channel，如果在go routine中使用，一定要做适当的延时，否则会输出来不及
	// 因为main已经跑完了，所以要延时一会，等待go routine
	// 有缓冲channel，在go routine里面的例子
	done := make(chan bool, 1)
	go func() {
		println("我可能会输出哦")
		done <- true
		println("我也可能会输出哦")
		<-done
		println("别注释done<-true哦，否则我输出不了")
	}()
	time.Sleep(time.Second * 1)
	println("完成")
}

// 多channel模式
func getMsgChannel(msg string, delay time.Duration) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Millisecond * delay)
		}
	}()
	return c
}

func test16() {
	c1 := getMsgChannel("第一", 600)
	c2 := getMsgChannel("第二", 500)
	c3 := getMsgChannel("第三", 5000)
	// 层次限制阻塞
	// c1 会阻塞 c2， c2 会阻塞 c3
	for i := 1; i <= 3; i++ {
		println(<-c1) // 除非c1有输入值，否则就阻塞下面的c2,c3
		println(<-c2) // 除非c2有输入值，否则就阻塞下面的c3
		println(<-c3) // 除非c3有输入值，否则就阻塞下一轮循环
	}
}

func test17() {
	c1 := getMsgChannel("第一", 600)
	c2 := getMsgChannel("第二", 500)
	c3 := getMsgChannel("第三", 5000)
	// select总是会把最先完成输入的channel输出，而且，互不限制
	// c1,c2,c3每两个互不限制
	for i := 1; i <= 9; i++ {
		select {
		case msg := <-c1:
			println(msg)
		case msg := <-c2:
			println(msg)
		case msg := <-c3:
			println(msg)
		}
	}
	/**
	 * 这个程序的运行结果：
	 *    第二 1，第三 1，第一 1，第二 2，第一 2，第二 3，第一 3，第三 2，第三 3
	 */
	/** 分析：前3次输出，“第一”，“第二”，“第三”，都有，而且
	 *  是随机顺序输出，因为协程的调度，第4，5，6次，由于“第二”只延时 500ms，
	 *  比 600ms 和 5000ms 都要小，那么它先输出，然后是“第一”，此时“第三”还不能输出，
	 *  因为它还在等5秒。此时已经输出5次，再过 500ms，"第三"的5秒还没走完，所以继续输出"第一"，
	 *  再过 100ms，500+100=600，"第二"也再完成了一次，那么输出。至此，"第一"和"第二"已经
	 *  把管道的 3 个值全部输出，9-7 = 2，剩下两个是 "第三"。此时，距离首次的 5000ms 完成，
	 *  还有，500-600-600 = 3800ms，达到后，"第三" 将输出，再过5秒，最后一次"第三输出"
	 */
}
