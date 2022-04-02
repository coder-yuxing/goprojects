package main

func main() {
	// 无缓冲 channel
	//ch1 := make(chan int)
	// 对同一个无缓冲 channel 只有对它进行接收操作的Goroutine 和对它进行发送操作的Goroutine都存在
	// 的情况下，通信才能进行，否则单方面的操作都会让对应Goroutine陷入挂起状态
	// 因此，对于无缓冲 Channel类型发送与接收操作，一定要放在两个不同的 Goroutine中记性，否则会导致deadlock
	//ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
	//n := <- ch1
	//println(n)

	ch2 := make(chan int, 1)
	n := <- ch2
	println(n)
	ch3 := make(chan int, 1)
	ch3 <- 17
	ch3 <- 23
}
