package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(1)
	fmt.Printf("list: %v", l)

	l.Front()
}
