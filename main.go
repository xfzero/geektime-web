package main //包声明

import "fmt" //方法关键字
func main() {
	print("Hello, Go!") //输出语句

	ForR()
}

func ForR() {
	i := 0
	for i < 4 {
		i++
		fmt.Printf("i:%d\n", i)
	}
	fmt.Printf("endi:%d\n", i)
}
