package main

import "fmt"

/*
	title::合成复用原则 Composite Reuse Principle
	继承和组合优先使用组合。
*/

type Gopher struct {
}

// GopherRaco 通过继承的方式 继承吃饭方法 新添加一个睡觉方法
type GopherRaco struct {
	Gopher
}

// GopherLori 通过组合的方式 继承吃饭的方法 并新添加一个睡觉的方法
type GopherLori struct {
	Gopher *Gopher
}

func (g *Gopher) Eat() {
	fmt.Println("囊鼠吃饭")
}

func (gl *GopherLori) Sleep() {
	fmt.Println("囊鼠Lori睡觉")
}

func (gr *GopherRaco) Sleep() {
	fmt.Println("囊鼠Raco睡觉")
}

func NewLori(gopher *Gopher) *GopherLori {
	return &GopherLori{Gopher: gopher}
}

func main() {
	// 继承方式继承
	Raco := new(GopherRaco)
	Raco.Eat()
	Raco.Sleep()

	// 组合方式继承
	Lori := NewLori(new(Gopher))
	Lori.Gopher.Eat()
	Lori.Sleep()
}
