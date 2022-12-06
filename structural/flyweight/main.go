package main

import "fmt"

// Flyweight 享元接口
type Flyweight interface {
	Operation(unshared UnsharedConcreteFlyweight)
}

// UnsharedConcreteFlyweight 非享元角色
type UnsharedConcreteFlyweight struct {
	info string
}

func (u *UnsharedConcreteFlyweight) SetInfo(info string) {
	u.info = info
}

func (u *UnsharedConcreteFlyweight) GetInfo() string {
	return u.info
}

// ConcreteFlyweight 具体享元
type ConcreteFlyweight struct {
	Key string
}

func (cf *ConcreteFlyweight) Operation(unshared UnsharedConcreteFlyweight) {
	fmt.Println("获取享元: " + cf.Key)
	fmt.Println("非分享信息: " + unshared.GetInfo())
}

// FlyweightFactory 享元工厂
type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func (factory *FlyweightFactory) GetFlyweight(key string) Flyweight {
	// 没有添加至工厂
	if _, ok := factory.flyweights[key]; !ok {
		fmt.Println("创建一个新的享元:", key)
		factory.flyweights[key] = &ConcreteFlyweight{Key: key}
	}
	return factory.flyweights[key]
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{flyweights: make(map[string]Flyweight, 0)}
}

func main() {
	factory := NewFlyweightFactory()

	f1 := factory.GetFlyweight("A")
	f2 := factory.GetFlyweight("A")
	f3 := factory.GetFlyweight("A")
	f4 := factory.GetFlyweight("B")
	f5 := factory.GetFlyweight("B")

	f1.Operation(UnsharedConcreteFlyweight{info: "first allocate A"})
	f2.Operation(UnsharedConcreteFlyweight{info: "second allocate A"})
	f3.Operation(UnsharedConcreteFlyweight{info: "third allocate A"})
	f4.Operation(UnsharedConcreteFlyweight{info: "first allocate B"})
	f5.Operation(UnsharedConcreteFlyweight{info: "second allocate B"})

}
