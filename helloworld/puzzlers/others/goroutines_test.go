package main

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func spawn(f func() error) <-chan error {
	c := make(chan error)

	go func() {
		c <- f()
	}()

	return c
}

func TestChannel(t *testing.T) {
	c := spawn(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("timeout")
	})
	fmt.Println(<-c)
}

func TestGoroutines(t *testing.T) {
	go fmt.Println("I am a goroutine")

	var c = make(chan int)
	go func(a, b int) {
		c <- a + b
	}(3, 4)

	var ch chan int

	ch = make(chan int)
	ch <- 1 + 1

}

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func TestLimitChannel(t *testing.T) {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}
