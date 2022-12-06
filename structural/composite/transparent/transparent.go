package main

import "fmt"

type Component interface {
	Add(c Component)
	Remove(c Component)
	GetChild(i int) Component
	Operation()
}

type Leaf struct {
	name string
}

func (l *Leaf) SetName(name string) {
	l.name = name
}

// 叶子节点的方法全部空实现

func (l *Leaf) Add(_ Component) {

}

func (l *Leaf) Remove(_ Component) {

}

func (l *Leaf) GetChild(_ int) Component {
	return nil
}

func (l *Leaf) Operation() {
	fmt.Println("Leaf " + l.name + " visited!")
}

type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{children: make([]Component, 0)}
}

func (cs *Composite) Add(c Component) {
	cs.children = append(cs.children, c)
}

func (cs *Composite) Remove(c Component) {
	flag := false
	switch c.(type) {
	case *Leaf:
		children := cs.children
		for i, v := range children {
			if value, ok := v.(*Leaf); ok {
				if value == c {
					flag = true
					fmt.Println("c is a leaf,found!")
					switch i {
					case 0:
						cs.children = children[1:]
					case len(cs.children) - 1:
						cs.children = children[0 : len(children)-1]
					default:
						cs.children = append(children[:i], children[i+1:]...)
					}
					break
				}

			}
		}
	case *Composite:
		children := cs.children
		for i, v := range children {
			if value, ok := v.(*Composite); ok {
				if value == c {
					flag = true
					fmt.Println("c is a composite,found!")
					switch i {
					case 0:
						cs.children = children[1:]
					case len(cs.children) - 1:
						cs.children = children[0 : len(children)-1]
					default:
						cs.children = append(children[:i], children[i+1:]...)
					}
					break
				}

			}
		}
	default:
		fmt.Println("c is a component")
	}

	if !flag {
		fmt.Println("c is not found!")
	}

}

func (cs *Composite) GetChild(i int) Component {
	return cs.children[i]
}

func (cs *Composite) Operation() {
	for _, child := range cs.children {
		child.Operation()
	}
}

func main() {
	fmt.Println("transparent pattern")
	// composite
	c0 := NewComposite()
	c1 := NewComposite()
	c2 := NewComposite()
	// leaf
	l1 := Leaf{name: "1"}
	l2 := Leaf{name: "2"}
	l3 := Leaf{name: "3"}
	l4 := Leaf{name: "4"}
	l5 := Leaf{name: "5"}
	l6 := Leaf{name: "6"}
	// 添加节点
	c0.Add(&l1)
	c0.Add(c1)
	c0.Add(&l5)
	c0.Add(&l6)
	c1.Add(&l2)
	c1.Add(c2)
	c2.Add(&l3)
	c2.Add(&l4)
	c0.Operation()

	fmt.Println("remove.......")
	// 移除树叶结构
	c0.Remove(&l1)
	// 移除树枝结构
	c0.Remove(c1)
	fmt.Println("after removed:")
	c0.Operation()
}
