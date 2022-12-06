package main

import "fmt"

// ProjectManager 指挥者::项目经理
type ProjectManager struct {
	decorator Decorator
}

func (pm *ProjectManager) SetDecorator(decorator Decorator) {
	pm.decorator = decorator
}

func (pm *ProjectManager) Decorate() *Parlour {
	decorator := pm.decorator
	decorator.BuildWall()
	decorator.BuildTV()
	decorator.BuildSofa()
	return decorator.GetResult()
}

// Decorator 建造者::抽象装修者
type Decorator interface {
	BuildWall()
	BuildTV()
	BuildSofa()
	GetResult() *Parlour
}

// Parlour 产品::客厅
type Parlour struct {
	wall string
	tv   string
	sofa string
}

func (p *Parlour) SetWall(wall string) {
	p.wall = wall
}

func (p *Parlour) SetTV(tv string) {
	p.tv = tv
}

func (p *Parlour) SetSofa(sofa string) {
	p.sofa = sofa
}

func (p *Parlour) Show() {
	fmt.Printf("wall:%s, TV:%s, sofa:%s\n", p.wall, p.tv, p.sofa)
}

// ConcreteDecorator1 具体装修者1
type ConcreteDecorator1 struct {
	parlour *Parlour
}

func (cd1 *ConcreteDecorator1) SetParlour(parlour *Parlour) {
	cd1.parlour = parlour
	fmt.Println("Set the parlour")
}

func (cd1 *ConcreteDecorator1) BuildWall() {
	cd1.parlour.SetWall("yellow wall")
	fmt.Println("Wall built by cd1!")
}

func (cd1 *ConcreteDecorator1) BuildTV() {
	cd1.parlour.SetTV("big TV")
	fmt.Println("TV built by cd1!")
}

func (cd1 *ConcreteDecorator1) BuildSofa() {
	cd1.parlour.SetSofa("soft Sofa")
	fmt.Println("Sofa built by cd1!")
}

func (cd1 *ConcreteDecorator1) GetResult() *Parlour {
	return cd1.parlour
}

// ConcreteDecorator2 具体装修者2
type ConcreteDecorator2 struct {
	parlour *Parlour
}

func (cd2 *ConcreteDecorator2) SetParlour(parlour *Parlour) {
	cd2.parlour = parlour
	fmt.Println("Set the parlour")
}

func (cd2 *ConcreteDecorator2) BuildWall() {
	cd2.parlour.SetWall("purple wall")
	fmt.Println("Wall built by cd2!")
}

func (cd2 *ConcreteDecorator2) BuildTV() {
	cd2.parlour.SetTV("small TV")
	fmt.Println("TV built by cd2!")
}

func (cd2 *ConcreteDecorator2) BuildSofa() {
	cd2.parlour.SetSofa("big soft Sofa")
	fmt.Println("Sofa built by cd2!")
}

func (cd2 *ConcreteDecorator2) GetResult() *Parlour {
	return cd2.parlour
}

func main() {
	fmt.Println("Manager ask decorator1 to decorate the parlour:")
	decorate1 := &ConcreteDecorator1{parlour: new(Parlour)}
	manager := new(ProjectManager)
	manager.SetDecorator(decorate1)
	p := manager.Decorate()
	p.Show()

	fmt.Println("Manager ask decorator2 to decorate the parlour:")
	decorate2 := &ConcreteDecorator2{parlour: new(Parlour)}
	manager.SetDecorator(decorate2)
	p2 := manager.Decorate()
	p2.Show()

}
