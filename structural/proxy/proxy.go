package main

import "fmt"

type Goods struct {
	Kind string // 商品种类
	Fact bool   // 商品真伪
}

// 抽象层

type Shopping interface {
	Buy(goods *Goods)
}

// 实现层

// KoreaShopping 韩国购物
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物, 买了 ", goods.Kind)
}

// AmericanShopping 美国购物
type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物, 买了 ", goods.Kind)
}

// BritainShopping 英国购物
type BritainShopping struct{}

func (bs *BritainShopping) Buy(goods *Goods) {
	fmt.Println("去英国进行了购物, 买了 ", goods.Kind)
}

// OverseasProxy 海外的代理
type OverseasProxy struct {
	shopping Shopping //代理某个主题，这里是抽象类型
}

func (op *OverseasProxy) Buy(goods *Goods) {
	// 1.验货
	if op.distinguish(goods) {
		// 2.购买-实现原有逻辑
		op.shopping.Buy(goods)
		// 3.海关检验
		op.check(goods)
	}

}

// distinguish 验货流程 pre
func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买。")
	}
	return goods.Fact
}

// check 海关安检
func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}

// NewProxy 创建一个代理
func NewProxy(shopping Shopping) *OverseasProxy {
	return &OverseasProxy{shopping: shopping}
}

func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	fmt.Println("---------不使用 代理模式-------")
	// 去韩国代购
	var shopping Shopping
	shopping = new(KoreaShopping)

	// 1-先验货
	if g1.Fact == true {
		fmt.Println("对[", g1.Kind, "]进行了辨别真伪.")
		//2-去韩国购买
		shopping.Buy(&g1)
		//3-海关安检
		fmt.Println("对[", g1.Kind, "] 进行了海关检查， 成功的带回祖国")
	}

	fmt.Println("---------使用 代理模式-------")
	proxy := NewProxy(shopping)
	proxy.Buy(&g1)
	proxy.Buy(&g2)
}
