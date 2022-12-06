package main

import "fmt"

type Cloneable interface {
	Clone() *SunWuKong
}

type SunWuKong struct {
	real   bool
	kind   string
	skills []string
}

func (s *SunWuKong) Clone() *SunWuKong {
	return &SunWuKong{
		real:   false,
		kind:   s.kind,
		skills: s.skills,
	}
}

func main() {
	king := &SunWuKong{real: true, kind: "monkey", skills: []string{"jump", "fly", "invisible"}}
	fake := king.Clone()
	fmt.Printf("%+v\n", king)
	fmt.Printf("%+v\n", fake)
	_, ok := interface{}(king).(Cloneable)
	fmt.Println(ok)

}
