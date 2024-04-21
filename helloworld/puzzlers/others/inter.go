package main

import (
	"errors"
	"io"
)

type MyInterface interface {
	M1() string
	M2(io.Writer, ...string)
}

type Interface1 interface {
	M1()
}
type Interface2 interface {
	M1(string)
	M2()
}

// Interface1 & Interface2 方法存在交集，此时要求重复的方法签名要保持一致
//type Interface3 interface{
//	Interface1
//	Interface2 // 编译器报错：duplicate method M1
//	M3()
//}

type QuackableAnimal interface {
	Quack()
}

type Duck struct{}

func (Duck) Quack() {
	println("duck quack!")
}

type Dog struct{}

func (Dog) Quack() {
	println("dog quack!")
}

type Bird struct{}

func (Bird) Quack() {
	println("bird quack!")
}

func AnimalQuackInForest(a QuackableAnimal) {
	a.Quack()
}

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func printNilInterface() {
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
}

func main() {
	animals := []QuackableAnimal{new(Duck), new(Dog), new(Bird)}
	for _, animal := range animals {
		AnimalQuackInForest(animal)
	}

	//err := returnsError()
	//if err != nil {
	//	fmt.Printf("error occur: %+v\n", err)
	//	return
	//}
	//fmt.Println("ok")

	printNilInterface()

	println(returnsError())

}
