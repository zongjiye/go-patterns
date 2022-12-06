package main

import "fmt"

type Programmer struct {
}

// 一个程序员 可能 要会前端、后端、运维、测试等等不同工种
// 这种平铺式设计在方法少的时候是可以这么做的，但如果所有的方法都集中在一个
// Programmer中会导致，添加新的方法时会出现问题导致之前的方法不可用

func (p *Programmer) DoBackend() {
	fmt.Println("正在进行后端开发...")
}

func (p *Programmer) DoFrontend() {
	fmt.Println("正在进行前端开发...")
}

func (p *Programmer) DoTest() {
	fmt.Println("正在进行测试...")
}

func (p *Programmer) DoOpt() {
	fmt.Println("正在进行运维...")
}

// 后端又分为多种语言 包括 Go、Python、Java、PHP、
// 前端又分为 Vue、React
// 如果都平铺添加的话 可能复用到以前的方法 导致方法不可用

func main() {
	p := new(Programmer)
	p.DoBackend()
	p.DoFrontend()
	p.DoTest()
	p.DoOpt()
}
