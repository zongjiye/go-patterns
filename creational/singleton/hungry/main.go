package main

import "fmt"

/*
饿汉式::含义是，在初始化单例唯一指针的时候，就已经提前开辟好了一个对象，申请了内存
*/
type singleton struct{}

var instance = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
