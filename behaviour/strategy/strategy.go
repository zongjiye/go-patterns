package main

import "fmt"

// SellStrategy 抽象策略接口
type SellStrategy interface {
	GetPrice(price float64) float64
}

// Discount 具体策略 打折
type Discount struct {
	off float64
}

func (d *Discount) GetPrice(price float64) float64 {
	fmt.Printf("执行策略A, 所有商品打折%.2f\n", d.off)
	return price * d.off
}

// Rebate 具体策略 返现
type Rebate struct {
	amount float64
}

func (r *Rebate) GetPrice(price float64) float64 {
	fmt.Printf("执行策略B, 所有商品满200 减%.2f\n", r.amount)

	if price >= 200 {
		price -= r.amount
	}

	return price
}

// Goods 环境类
type Goods struct {
	Origin   float64
	strategy SellStrategy
}

func (g *Goods) SetStrategy(strategy SellStrategy) {
	g.strategy = strategy
}

func (g *Goods) GetStrategy() SellStrategy {
	return g.strategy
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("原价值 ", g.Origin, " .")
	return g.strategy.GetPrice(g.Origin)
}

func main() {
	nike := Goods{
		Origin: 200.0,
	}
	//上午 ，商场执行策略A
	nike.SetStrategy(&Discount{off: 0.8})
	fmt.Println("上午nike鞋卖", nike.SellPrice())

	//下午， 商场执行策略B
	nike.SetStrategy(&Rebate{amount: 100})
	fmt.Println("下午nike鞋卖", nike.SellPrice())
}
