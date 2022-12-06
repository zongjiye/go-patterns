package main

import "fmt"

type Hero interface {
	Desc()
}

type Galen struct {
}

type Yasuo struct {
}

type Yongen struct {
}

func (g *Galen) Desc() {
	fmt.Println("I'm Galen")
}

func (y *Yasuo) Desc() {
	fmt.Println("I'm Yasuo")
}

func (y *Yongen) Desc() {
	fmt.Println("I'm Yongen")
}

type HeroFactory struct {
}

func (hf *HeroFactory) CreateHero(name string) (hero Hero) {
	switch name {
	case "Galen":
		hero = new(Galen)
	case "Yasuo":
		hero = new(Yasuo)
	case "Yongen":
		hero = new(Yongen)
	default:
		panic("To be realized")
	}
	return
}

func main() {
	factory := new(HeroFactory)
	// 创建亚索
	yasuo := factory.CreateHero("Yasuo")
	yasuo.Desc()
	// 创建永恩
	yongen := factory.CreateHero("Yongen")
	yongen.Desc()
	// 创建盖伦
	galen := factory.CreateHero("Galen")
	galen.Desc()
}
