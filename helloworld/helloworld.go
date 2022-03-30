package main

import "fmt"

func main() {
   fmt.Println("hello, world")

   test := Test{Name: "test", Numbers: []int{1, 2, 3}}
   fmt.Println(test.Name)
   fmt.Println(test.Numbers[2])
   rename(test)
   fmt.Println(test.Name)
   fmt.Println(test.Numbers[2])
   rename1(&test)
   fmt.Println(test.Name)
   fmt.Println(test.Numbers[2])

   // Go方法的实质是以方法的 receiver 参数作为第一个参数的普通函数
   // Go方法/函数传参采用值传递
   // 故当方法的receiver参数为 T时，实际传入的是 T类型实例的副本，因此方法内对该参数的任何修改都对原实例没有影响
   // 当方法receiver参数为 *T时，实际传入的是T类型示实例的地址，因此方法内对该参数的操作会反映到原T类型实例上
   var t T
   println(t.a) // 0

   t.M1()
   println(t.a) // 0

   p := &t
   p.M2()
   println(t.a) // 11
}

type Test struct {
   Name string
   Numbers []int
}

func rename(t Test) {
   t.Name = "rTest"
   t.Numbers[2] = 4
}

func rename1(t *Test) {
   t.Name = "rTest"
   t.Numbers[2] = 5
}