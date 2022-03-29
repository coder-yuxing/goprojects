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