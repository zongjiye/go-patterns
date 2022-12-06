package main

import "fmt"

// DrinkMaker 抽象类，制作饮料,包裹一个模板的全部实现步骤
type DrinkMaker interface {
	// PourWater 倒水 特殊方法公用
	PourWater()

	// BoilWater Brew PourInCup AddThings WantAddThings抽象方法子类可重写
	BoilWater()          // 烧开水
	Brew()               // 冲泡
	PourInCup()          // 倒入杯中
	AddThings()          // 添加酌料
	WantAddThings() bool // 是否加入酌料

	// MakeBeverage 模板方法:: 制作茶饮
	MakeBeverage()
}

// Maker 模板类 实现定义的抽象类
type Maker struct {
	DrinkMaker DrinkMaker
}

func NewMaker(maker DrinkMaker) *Maker {
	return &Maker{DrinkMaker: maker}
}

func (m *Maker) PourWater() {
	fmt.Println("将水倒入到烧水壶中")
}

func (m *Maker) BoilWater() {
	if m.DrinkMaker == nil {
		return
	}
	m.DrinkMaker.BoilWater()
}

func (m *Maker) Brew() {
	if m.DrinkMaker == nil {
		return
	}
	m.DrinkMaker.Brew()
}

func (m *Maker) PourInCup() {
	if m.DrinkMaker == nil {
		return
	}
	m.DrinkMaker.PourInCup()
}

func (m *Maker) AddThings() {
	if m.DrinkMaker == nil {
		return
	}
	m.DrinkMaker.AddThings()
}

func (m *Maker) WantAddThings() bool {
	if m.DrinkMaker != nil {
		return m.DrinkMaker.WantAddThings()
	}
	return false
}

// MakeBeverage 封装模板方法
func (m *Maker) MakeBeverage() {
	m.PourWater()
	m.BoilWater()
	m.Brew()
	m.PourInCup()
	if m.WantAddThings() {
		m.AddThings()
	}

}

// CoffeeMaker 具体的模板子类 制作咖啡 继承模板
type CoffeeMaker struct {
	Maker
}

func (cm *CoffeeMaker) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

func (cm *CoffeeMaker) Brew() {
	fmt.Println("用水冲咖啡豆")
}

func (cm *CoffeeMaker) PourInCup() {
	fmt.Println("将冲好的咖啡倒入陶瓷杯中")
}

func (cm *CoffeeMaker) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (cm *CoffeeMaker) WantAddThings() bool {
	return true
}

// TeaMaker 具体的模板子类 制作茶 继承模板
type TeaMaker struct {
	Maker
}

func (tm *TeaMaker) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

func (tm *TeaMaker) Brew() {
	fmt.Println("用水冲茶叶")
}

func (tm *TeaMaker) PourInCup() {
	fmt.Println("将泡好的茶倒入茶壶中")
}

func main() {
	// 制作咖啡
	coffeeMaker := NewMaker(new(CoffeeMaker))
	coffeeMaker.MakeBeverage()
	fmt.Println("--------------")
	// 制作茶
	teaMaker := NewMaker(new(TeaMaker))
	teaMaker.MakeBeverage()
}
