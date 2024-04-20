package main

import "fmt"

var container = []string{"zero", "one", "two"}

type MyString = string
type MyString2 string

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	//fmt.Printf("The element is %q.\n", container[1])

	// 类型断言表达式 x.(T) x必须为接口类型
	if value, ok := interface{}(container).([]string); ok {
		fmt.Printf("The element is %q.\n", value[1])
	} else {
		fmt.Println("not ok")
	}

	v := interface{}(container).(map[int]string)
	fmt.Printf("The element is %q.\n", v[1])

	fmt.Println(string(-1))

	var s MyString2 = "hello"
	var s1 = string(s)
	fmt.Println(s1)

}
