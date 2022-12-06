package abstract

import "fmt"

// Nvidia厂商

type NvidiaCpu struct {
}

type NvidiaGpu struct {
}

type NvidiaMemory struct {
}

func (nc *NvidiaCpu) Calculate() {
	fmt.Println("英伟达Cpu")
}

func (ng *NvidiaGpu) Display() {
	fmt.Println("英伟达Gpu")
}

func (nm *NvidiaMemory) Storage() {
	fmt.Println("英伟达Memory")
}

// NvidiaFactory 英伟达工厂
type NvidiaFactory struct {
}

func (nf *NvidiaFactory) CreateCpu() Cpu {
	return new(NvidiaCpu)
}

func (nf *NvidiaFactory) CreateGpu() Gpu {
	return new(NvidiaGpu)
}

func (nf *NvidiaFactory) CreateMemory() Memory {
	return new(NvidiaMemory)
}
