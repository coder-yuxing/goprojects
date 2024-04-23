package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 无缓冲channel的惯用法
// 1.用作信号传递
type signal struct {
}

func worker() {
	println("worker is working...")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		f()
		c <- signal{}
	}()
	return c
}

func TestChannelSignal(t *testing.T) {
	println("start a worker...")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")
}

// 2. 1对n的信号通知机制
func worker2(i int) {
	fmt.Printf("worker %d is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d work done!\n", i)
}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			// 等待信号，因为是无缓冲channel, 所以这里会阻塞
			// 因此所有worker运行到此都会阻塞
			<-groupSignal
			fmt.Printf("worker %d: start to work\n", i)
			defer wg.Done()
			f(i)
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()

	return c
}

func TestChannelSignal2(t *testing.T) {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(worker2, 5, groupSignal)
	time.Sleep(5 * time.Second)
	// 全部worker都阻塞在起始位置，此处关闭groupSignal, 相当于通知所有worker工作开始工作
	close(groupSignal)
	// 等待所有worker工作完成
	<-c
	fmt.Println("all workers work done!")
}

type counter struct {
	sync.Mutex
	i int
}

var cter counter

func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func TestCounter(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, Increase())
		}(i)
	}
	wg.Wait()
}

// 使用无缓存channel实现计数器
type chanCounter struct {
	c chan int
	i int
}

func NewChanCounter() *chanCounter {
	cter := &chanCounter{
		c: make(chan int),
	}

	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cc *chanCounter) Increase() int {
	return <-cc.c
}

func TestChanCounter(t *testing.T) {
	cter := NewChanCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, cter.Increase())
		}(i)
	}
	wg.Wait()
}
