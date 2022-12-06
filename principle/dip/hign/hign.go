package main

import "fmt"

// 一个高耦合的例子
// 人使用电脑 这是一个排列组合

type Lenovo struct {
}

type HP struct {
}

type MSI struct {
}

func (l *Lenovo) Run() {
	fmt.Println("Lenovo is running ...")
}

func (msi *MSI) Run() {
	fmt.Println("MSI is running...")
}

func (hp *HP) Run() {
	fmt.Println("HP is running...")
}

type Mark struct {
}

type Viper struct {
}

// Mark use HP
// Mark use MSI

func (m *Mark) UseHP(hp *HP) {
	fmt.Println("Mark use hp")
	hp.Run()
}

func (m *Mark) UseMSI(msi *MSI) {
	fmt.Println("Mark use msi")
	msi.Run()
}

// Viper use Lenovo
// Viper use MSI

func (v *Viper) UseLenovo(lenovo *Lenovo) {
	fmt.Println("Viper use lenovo")
	lenovo.Run()
}

func (v *Viper) UseMSI(msi *MSI) {
	fmt.Println("Viper use msi")
	msi.Run()
}

func main() {
	hp := new(HP)
	msi := new(MSI)
	lenovo := new(Lenovo)

	// Mark use HP
	// Mark use MSI
	mark := new(Mark)
	mark.UseHP(hp)
	mark.UseMSI(msi)

	// Viper use Lenovo
	// Viper use MSI
	viper := new(Viper)
	viper.UseLenovo(lenovo)
	viper.UseMSI(msi)
}
