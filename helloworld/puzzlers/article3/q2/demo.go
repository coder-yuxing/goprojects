package main

import (
	"flag"
	"fmt"
	"github.com/yuxing/goprojects/helloworld/puzzlers/article3/q2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

var s string = "package"

func main() {
	flag.Parse()
	fmt.Println(lib.Hello(name))

	fmt.Println("s=", s)

	var s string = "你好"
	fmt.Println("s=", s)
	{
		//var s string = "世界"
		//s = "世界"
		s := 1
		fmt.Println("s=", s)
	}
	fmt.Println("s=", s)
}
