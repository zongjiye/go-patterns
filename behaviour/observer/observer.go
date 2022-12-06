package main

import "fmt"

// Listener 抽象观察者
type Listener interface {
	OnTeacherComing() // 观察者检测到变动时, 出发的操作
}

// Notifier 被通知的目标
type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// 实现层

type StuZhang3 struct {
	BadThing string
}

func (s *StuZhang3) OnTeacherComing() {
	fmt.Println("张3 停止 ", s.BadThing)
}

func (s *StuZhang3) DoBadThing() {
	fmt.Println("张3 正在", s.BadThing)
}

type StuZhao4 struct {
	BadThing string
}

func (s *StuZhao4) OnTeacherComing() {
	fmt.Println("赵4 停止 ", s.BadThing)
}

func (s *StuZhao4) DoBadThing() {
	fmt.Println("赵4 正在", s.BadThing)
}

type StuWang5 struct {
	BadThing string
}

func (s *StuWang5) OnTeacherComing() {
	fmt.Println("王5 停止 ", s.BadThing)
}

func (s *StuWang5) DoBadThing() {
	fmt.Println("王5 正在", s.BadThing)
}

// ClassMonitor 通知者 实现Notifier
type ClassMonitor struct {
	listenerList []Listener
}

func (c *ClassMonitor) AddListener(listener Listener) {
	c.listenerList = append(c.listenerList, listener)
}

func (c *ClassMonitor) RemoveListener(listener Listener) {
	for i, l := range c.listenerList {
		if listener == l {
			c.listenerList = append(c.listenerList[:i], c.listenerList[i+1:]...)
			break
		}
	}
}

func (c *ClassMonitor) Notify() {
	for _, l := range c.listenerList {
		l.OnTeacherComing()
	}
}

func main() {
	s1 := &StuZhang3{
		BadThing: "抄作业",
	}
	s2 := &StuZhao4{
		BadThing: "玩王者荣耀",
	}
	s3 := &StuWang5{
		BadThing: "看赵四玩王者荣耀",
	}

	// 观察者
	classMonitor := new(ClassMonitor)

	fmt.Println("上课了，但是老师没有来，学生们都在忙自己的事...")
	s1.DoBadThing()
	s2.DoBadThing()
	s3.DoBadThing()

	classMonitor.AddListener(s1)
	classMonitor.AddListener(s2)
	classMonitor.AddListener(s3)

	fmt.Println("这时候老师来了，班长给学什么使了一个眼神...")
	classMonitor.Notify()
}
