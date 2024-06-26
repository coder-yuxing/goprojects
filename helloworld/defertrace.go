package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// 使用defer跟踪函数的执行过程
var mu sync.Mutex
var m = make(map[uint64]int)

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	gid := curGoroutineID()

	mu.Lock()
	indents := m[gid]
	m[gid] = indents + 1
	mu.Unlock()
	printTrace(gid, name, "->", indents+1)
	return func() {
		mu.Lock()
		indents := m[gid]
		m[gid] = indents - 1
		mu.Unlock()
		printTrace(gid, name, "<-", indents)
	}
}

// trace2/trace.go
var goroutineSpace = []byte("goroutine ")

// 获取GoroutineID
func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q3", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q3: %v", b, err))
	}
	return n
}

func printTrace(id uint64, name, arrow string, ident int) {
	indents := ""
	for i := 0; i < ident; i++ {
		indents += "   "
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

func main2() {
	defer Trace()()
	foo()
}
