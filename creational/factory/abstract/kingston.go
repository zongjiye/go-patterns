package abstract

import "fmt"

// Kingston厂商

type KingstonCpu struct {
}

type KingstonGpu struct {
}

type KingstonMemory struct {
}

func (kc *KingstonCpu) Calculate() {
	fmt.Println("Kingston Cpu")
}

func (kg *KingstonGpu) Display() {
	fmt.Println("Kingston Gpu")
}

func (km *KingstonMemory) Storage() {
	fmt.Println("Kingston Memory")
}

// KingstonFactory Kingston工厂
type KingstonFactory struct {
}

func (kf *KingstonFactory) CreateCpu() Cpu {
	return new(KingstonCpu)
}

func (kf *KingstonFactory) CreateGpu() Gpu {
	return new(KingstonGpu)
}

func (kf *KingstonFactory) CreateMemory() Memory {
	return new(KingstonMemory)
}
