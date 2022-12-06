package main

import (
	"patterns/creational/factory/abstract"
)

/*
	设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口，显卡具有显示功能（display，功能实现只要打印出意义即可），内存具有存储功能（storage），cpu具有计算功能（calculate）。
	现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
	要求组装两台电脑，

		1台（Intel的CPU，Intel的显卡，Intel的内存）
		1台（Intel的CPU， nvidia的显卡，Kingston的内存）

	用抽象工厂模式实现。
*/

type Computer struct {
	Cpu    abstract.Cpu
	Gpu    abstract.Gpu
	Memory abstract.Memory
}

func NewComputer(cpu abstract.Cpu, gpu abstract.Gpu, memory abstract.Memory) *Computer {
	return &Computer{
		Cpu:    cpu,
		Gpu:    gpu,
		Memory: memory,
	}
}

func main() {
	// 组装两台
	var intelFac abstract.Factory
	var nvidiaFac abstract.Factory
	var kingstonFac abstract.Factory
	intelFac = new(abstract.IntelFactory)
	nvidiaFac = new(abstract.NvidiaFactory)
	kingstonFac = new(abstract.KingstonFactory)

	// 第一台电脑
	c1 := NewComputer(intelFac.CreateCpu(), intelFac.CreateGpu(), intelFac.CreateMemory())
	c1.Cpu.Calculate()
	c1.Gpu.Display()
	c1.Memory.Storage()
	// 第二台电脑
	c2 := NewComputer(intelFac.CreateCpu(), nvidiaFac.CreateGpu(), kingstonFac.CreateMemory())
	c2.Cpu.Calculate()
	c2.Gpu.Display()
	c2.Memory.Storage()
}
