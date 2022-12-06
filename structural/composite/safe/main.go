package main

import "fmt"

type Component interface {
	Operation()
}

type Leaf struct {
	name string
}

func (l *Leaf) SetName(name string) {
	l.name = name
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

func (c2 *Composite) Add(c Component) {
	c2.children = append(c2.children, c)
}

func (c2 *Composite) Remove(c Component) {
	flag := false
	switch c.(type) {
	case *Leaf:
		children := c2.children
		for i, v := range children {
			if value, ok := v.(*Leaf); ok {
				if value == c {
					flag = true
					fmt.Println("c is a leaf, found!")
					switch i {
					case 0:
						c2.children = children[1:]
					case len(children) - 1:
						c2.children = children[:len(children)-1]
					default:
						c2.children = append(children[:i], children[i+1:]...)
					}
					break
				}
			} else {
				if composite, ok2 := v.(*Composite); ok2 {
					composite.Remove(c)
				}
			}
		}
	case *Composite:
		children := c2.children
		for i, v := range children {
			if value, ok := v.(*Composite); ok {
				if value == c {
					flag = true
					fmt.Println("c is a composite, found!")
					switch i {
					case 0:
						c2.children = children[1:]
					case len(children) - 1:
						c2.children = children[:len(children)-1]
					default:
						c2.children = append(children[:i], children[i+1:]...)
					}
					break
				} else {
					value.Remove(c)
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

func (c2 *Composite) GetChild(i int) Component {
	return c2.children[i]
}

func (c2 *Composite) Operation() {
	for _, child := range c2.children {
		child.Operation()
	}
}

func main() {
	// 树枝结构
	c0 := NewComposite()
	c1 := NewComposite()
	c2 := NewComposite()
	// 树叶结构
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
	// 遍历-移除节点前
	fmt.Println("before removed")
	c0.Operation()
	// 移除节点
	c0.Remove(c1)
	c0.Remove(&Leaf{name: "8"})
	c3 := NewComposite()
	c0.Remove(c3)
	// 遍历-移除节点后
	fmt.Println("after removed:")
	c0.Operation()
}
