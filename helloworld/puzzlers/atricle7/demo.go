package main

import "fmt"

func main() {

	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)

	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)

	s2 = append(s2, 1)
	fmt.Printf("The value of s2: %d\n", s2)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)

	s5 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s6 := s5[4:6]
	s6 = append(s6, 98)                     // 此时 s5 的值会被修改 [0 1 2 3 4 5 98 7 8]
	fmt.Printf("The value of s5: %d\n", s5) // [0 1 2 3 4 5 98 7 8]
	fmt.Printf("The value of s6: %d\n", s6) // [4 5 98]

}
