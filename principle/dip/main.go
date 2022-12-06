package main

import "fmt"

/*
	将模块分为3个层次，抽象层、实现层、业务逻辑层
	title::依赖倒转原则 Dependence  Inversion Principle
*/

// 抽象层:: 定义两个抽象接口

type Computer interface {
	Run()
}

type User interface {
	Use(computer Computer)
}

// 实现层

type Lenovo struct {
}

func (l *Lenovo) Run() {
	fmt.Println("Lenovo is running...")
}

type HP struct {
}

func (hp *HP) Run() {
	fmt.Println("HP is running...")
}

type MSI struct {
}

func (msi *MSI) Run() {
	fmt.Println("MSI is running...")
}

type Mark struct {
}

func (m *Mark) Use(computer Computer) {
	fmt.Println("Mark use computer")
	computer.Run()
}

type Viper struct {
}

func (v *Viper) Use(computer Computer) {
	fmt.Println("Viper use computer")
	computer.Run()
}

func main() {
	var computer Computer
	var user User

	// Mark Use msi
	// Mark Use hp
	user = new(Mark)
	computer = new(MSI)
	user.Use(computer)

	computer = new(HP)
	user.Use(computer)

	// Viper use Lenovo
	// Viper use msi
	user = new(Viper)
	computer = new(Lenovo)
	user.Use(computer)

	computer = new(MSI)
	user.Use(computer)
}
