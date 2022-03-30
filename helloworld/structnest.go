package main

import (
	"fmt"
	"io"
	"strings"
)

// 结构体嵌套

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}

type S struct {
	*MyInt
	t
	io.Reader
	s string
	n int
}

func test() {
	m := MyInt(17)
	r := strings.NewReader("hello, go")
	s := S{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}

	sl := make([]byte, len("hello, go"))
	//s.Reader.Read(sl)
	s.Read(sl)  // 组合中代理模式的体现
	fmt.Println(string(sl))
	//s.MyInt.Add(5)
	s.Add(5)
	fmt.Println(*(s.MyInt))
}

