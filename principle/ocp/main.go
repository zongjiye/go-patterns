package main

import "fmt"

type AbstractProgrammer interface {
	Do()
}

// 满足开闭原则的设计 Open-Closed Principle
// 定义抽象类型
// 对不同职责的员工 实现同一个相同的方法 实现解耦
// 当新添加一种 工种时 是需要 新添加代码即可
// 而不是修改原来的代码

type BackProgrammer struct {
}

type FrontProgrammer struct {
}

type TestProgrammer struct {
}

type OptProgrammer struct {
}

func (b *BackProgrammer) Do() {
	fmt.Println("正在进行后端开发......")
}

func (b *FrontProgrammer) Do() {
	fmt.Println("正在进行前端开发......")
}

func (b *TestProgrammer) Do() {
	fmt.Println("正在进行测试......")
}

func (b *OptProgrammer) Do() {
	fmt.Println("正在进行运维......")
}

func ProgrammerDo(programmer AbstractProgrammer) {
	programmer.Do()
}

func main() {
	bp := new(BackProgrammer)
	bp.Do()

	fp := new(FrontProgrammer)
	fp.Do()

	tp := new(TestProgrammer)
	tp.Do()

	op := new(OptProgrammer)
	op.Do()

	ProgrammerDo(bp)
	ProgrammerDo(fp)
}
