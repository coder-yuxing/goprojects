package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"syscall"
)

func main() {
	// 整个文件读入内存
	// 1. 指定文件名直接将数据读取入内存
	content, err := readAllFileContent("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	content, err = readAllFileContent1("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	// 2.先创建句柄在读取
	content, err = readAllFileContent2("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	// 2.1 可以直接使用 os.OpenFile 方法，与上述操作是等价的
	content, err = readAllFileContent3("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	fmt.Println("-------------------------------")
	readLine()

	fmt.Println("-------------------------------")
	readFixedLengthBytes()
	fmt.Println("-------------------------------")
	readBySyscall()
}


func readAllFileContent(filePath string) (string, error){
	if len(filePath) == 0 {
		return "", errors.New("filePath is empty")
	}
	content, err := os.ReadFile(filePath)
	return string(content), err
}

func readAllFileContent1(filePath string) (string, error) {
	if len(filePath) == 0 {
		return "", errors.New("filePath is empty")
	}
	content, err := ioutil.ReadFile(filePath)
	return string(content), err
}

func readAllFileContent2(filePath string) (string, error) {
	if len(filePath) == 0 {
		return "", errors.New("filePath is empty")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return string(content), err
}

func readAllFileContent3(filePath string) (string, error) {
	if len(filePath) == 0 {
		return "", errors.New("filePath is empty")
	}

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return string(content), err
}


// 单行读取
// 1. bufio.ReadLone() bufio源码注释中说明，ReadLine是低级库，不适合普通用户使用
// 2. bufio.ReadBytes('\n')
// 3. bufio.ReadString('\n')
func readLine() {
	file, err := os.Open("b.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		//lineBytes, err := r.ReadBytes('\n')
		lineBytes, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(string(lineBytes))
	}
}

// 每次仅读取一行数据，固然可以解决内存占用过大的问题，但需要注意的是，并不是所有文件都有换行符'\n'
// 因此，对于一些不换行的大文件，还需使用其他方式处理
func readFixedLengthBytes() {
	file, err := os.Open("b.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func readBySyscall() {
	// fc: file descriptor
	fd, err := syscall.Open("b.txt", syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	var wg sync.WaitGroup
	wg.Add(2)
	dataChan := make(chan []byte)
	go func() {
		wg.Done()
		for {
			data := make([]byte, 100)
			n, _ := syscall.Read(fd, data)
			if n == 0 {
				break
			}
			dataChan <- data
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case dataChan, ok := <- dataChan:
				if !ok {
					return
				}
				fmt.Println(string(dataChan))
			default:

			}
		}
	}()

	wg.Wait()
}
