package main

import "fmt"

type User struct {
	Name string
	Age  uint32
}

//结构体接收器；一般用于不希望被修改的结构体，只是复制变量
func (u User) ChangeName(newName string) {
	u.Name = newName
}

//指针接收器
func (u *User) ChangeName2(newName string) {
	u.Name = newName
}

func main() {
	newName := "Tony"
	user := &User{
		Name: "Jim",
	}
	user.ChangeName(newName)
	fmt.Printf("user:%v \n", user) //name依然是Jim
	user.ChangeName2(newName)
	fmt.Printf("user:%v \n", user) //name改成了Tony
}

type Handle func()

func (h Handle) Hello() {

}
