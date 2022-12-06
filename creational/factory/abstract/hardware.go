package abstract

// == 抽象层 ==

type Gpu interface {
	Display() // 显示功能
}

type Cpu interface {
	Calculate() // 计算功能
}

type Memory interface {
	Storage() // 存储功能
}

// Factory 硬件工厂
type Factory interface {
	CreateCpu() Cpu
	CreateGpu() Gpu
	CreateMemory() Memory
}
