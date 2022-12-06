package main

import "fmt"

// Charger 充电器接口
type Charger interface {
	Charge()
}

// PhoneCharger 手机充电器
type PhoneCharger struct {
}

// Charge5v 使用5v电压充电
func (pc *PhoneCharger) Charge5v() {
	fmt.Println("使用5V电压进行充电")
}

// ComputerCharger 电脑充电器
type ComputerCharger struct {
}

// Charge110v 使用110v电压充电
func (cc *ComputerCharger) Charge110v() {
	fmt.Println("使用110V电压进行充电")
}

// PhoneAdapter 手机充电适配器
type PhoneAdapter struct {
	charger PhoneCharger
}

func (pa *PhoneAdapter) new() {
	pa.charger = PhoneCharger{}
}

func (pa *PhoneAdapter) Charge() {
	pa.new()
	pa.charger.Charge5v()
}

// ComputerAdapter 电脑充电适配器
type ComputerAdapter struct {
	charger ComputerCharger
}

func (ca *ComputerAdapter) new() {
	ca.charger = ComputerCharger{}
}

func (ca *ComputerAdapter) Charge() {
	ca.new()
	ca.charger.Charge110v()
}

func main() {
	var charger Charger

	fmt.Println("使用手机充电器适配器")
	pAdapter := &PhoneAdapter{}
	charger = pAdapter
	charger.Charge()

	fmt.Println("使用电脑充电器适配器")
	cAdapter := &ComputerAdapter{}
	charger = cAdapter
	charger.Charge()
}
