package main

import "fmt"

/*
	title::单一职责原则 Single Responsibility Principle
	类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。
*/

type WorkComputer struct {
}

type GameComputer struct {
}

func (wc *WorkComputer) OnWork() {
	fmt.Println("正在工作中的computer")
}

func (gc *GameComputer) OnGame() {
	fmt.Println("正在游戏中的computer")
}

func main() {
	// 工作电脑 提供工作方法
	wc := new(WorkComputer)
	wc.OnWork()
	// 游戏电脑 提供游戏方法
	gc := new(GameComputer)
	gc.OnGame()
}
