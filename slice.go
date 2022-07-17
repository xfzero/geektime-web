package main

import "fmt"

func main() {
	test2()
}

func test1() {
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1:%v, len:%d, cap:%d \n", s1, len(s1), cap(s1))

	s2 := s1[:2]
	fmt.Printf("s2:%v, len:%d, cap:%d \n", s2, len(s2), cap(s2))

	s3 := s1[2:]
	fmt.Printf("s3:%v, len:%d, cap:%d \n", s3, len(s3), cap(s3))
	fmt.Printf("s1:%v, len:%d, cap:%d \n", s1, len(s1), cap(s1))
}

func test2() {
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1:%v, len:%d, cap:%d \n", s1, len(s1), cap(s1))

	s2 := s1[:2]
	fmt.Printf("s2:%v, len:%d, cap:%d \n", s2, len(s2), cap(s2))

	s3 := s1[2:]
	fmt.Printf("s3:%v, len:%d, cap:%d \n", s3, len(s3), cap(s3))
	fmt.Printf("s1:%v, len:%d, cap:%d \n", s1, len(s1), cap(s1))

	s2[1] = 5
	fmt.Printf("s1:%v, len:%d, cap:%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v, len:%d, cap:%d \n", s2, len(s2), cap(s2))

	s3[1] = 6
	fmt.Printf("s1:%v, len:%d, cap:%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s3:%v, len:%d, cap:%d \n", s3, len(s3), cap(s3))

	s3 = append(s3, 7)
	fmt.Printf("s1:%v, len:%d, cap:%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v, len:%d, cap:%d \n", s2, len(s2), cap(s2))
	fmt.Printf("s3:%v, len:%d, cap:%d \n", s3, len(s3), cap(s3))

	s3[1] = 8
	fmt.Printf("s1:%v, len:%d, cap:%d \n", s1, len(s1), cap(s1))
}
