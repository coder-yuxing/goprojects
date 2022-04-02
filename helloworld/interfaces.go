package main

import "fmt"

// 接口即契约

// 类型断言
func main3() {
	var a int64 = 13
	var i interface{} = a
	v1, ok := i.(int64)
	fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok)
	v2, ok := i.(string)
	fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok)
	v3 := i.(int64)
	fmt.Printf("v3=%s, the type of v3 is %T\n", v3, v3)
	v4 := i.(string)
	fmt.Printf("v4=%s, the type of v4 is %T\n", v4, v4)

}
