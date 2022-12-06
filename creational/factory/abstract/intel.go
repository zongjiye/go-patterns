package abstract

import "fmt"

// Intel厂商

type IntelCpu struct {
}

type IntelGpu struct {
}

type IntelMemory struct {
}

func (ic *IntelCpu) Calculate() {
	fmt.Println("英特尔Cpu")
}

func (ig *IntelGpu) Display() {
	fmt.Println("英特尔Gpu")
}

func (im *IntelMemory) Storage() {
	fmt.Println("英特尔Memory")
}

// IntelFactory 英特尔工厂
type IntelFactory struct {
}

func (f *IntelFactory) CreateCpu() Cpu {
	return new(IntelCpu)
}

func (f *IntelFactory) CreateGpu() Gpu {
	return new(IntelGpu)
}

func (f *IntelFactory) CreateMemory() Memory {
	return new(IntelMemory)
}
