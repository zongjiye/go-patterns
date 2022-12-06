package main

import "fmt"

type Phone interface {
	Show()
}

type XiaoMi struct {
	Spec string
	Name string
}

func (x *XiaoMi) Show() {
	fmt.Println("展示出我的小米手机")
}

type PhoneDecorator interface {
	Phone
}

type FilmDecorator struct {
	phone Phone
}

type ShellDecorator struct {
	phone Phone
}

func (fd *FilmDecorator) AddFilm() {
	fmt.Println("给手机贴膜")
}

func (sd *ShellDecorator) AddShell() {
	fmt.Println("给手机套壳")
}

func (fd *FilmDecorator) Show() {
	fd.AddFilm()
	fd.phone.Show()
}

func (sd *ShellDecorator) Show() {
	sd.AddShell()
	sd.phone.Show()
}

func main() {
	// 未经过装饰的小米手机
	xiaomi := &XiaoMi{Spec: "12G, 516GB", Name: "小米12sUltra"}
	xiaomi.Show()
	// 经过装饰的小米手机

	// 贴膜
	moXiaomi := &FilmDecorator{phone: xiaomi}
	moXiaomi.Show()
	// 带壳
	keXiaomi := &ShellDecorator{phone: xiaomi}
	keXiaomi.Show()
	// 贴膜 + 带壳
	kmXiaomi := &ShellDecorator{phone: moXiaomi}
	kmXiaomi.Show()
}
