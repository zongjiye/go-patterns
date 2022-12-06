package main

import "fmt"

type HeroicSkill interface {
	Desc()
}

type HeroicSkillFactory interface {
	CreateHeroicSkill() HeroicSkill
}

// SteelTempest 斩钢闪
type SteelTempest struct{}

// WindWall 风墙
type WindWall struct{}

// SweepingBlade 踏前斩
type SweepingBlade struct{}

// LastBreath 狂风绝息斩
type LastBreath struct{}

func (st *SteelTempest) Desc() {
	fmt.Println("亚索Q技能：斩钢闪")
}

func (ww *WindWall) Desc() {
	fmt.Println("亚索W技能：风墙")
}

func (sb *SweepingBlade) Desc() {
	fmt.Println("亚索E技能：踏前斩")
}

func (lb *LastBreath) Desc() {
	fmt.Println("亚索R技能：狂风绝息斩")
}

type SteelTempestFactory struct{}

type WindWallFactory struct{}

type SweepingBladeFactory struct{}

type LastBreathFactory struct{}

func (stf *SteelTempestFactory) CreateHeroicSkill() HeroicSkill {
	return new(SteelTempest)
}

func (stf *WindWallFactory) CreateHeroicSkill() HeroicSkill {
	return new(WindWall)
}

func (stf *SweepingBladeFactory) CreateHeroicSkill() HeroicSkill {
	return new(SweepingBlade)
}

func (stf *LastBreathFactory) CreateHeroicSkill() HeroicSkill {
	return new(LastBreath)
}

func main() {
	// 生产亚索的四个技能Q W E R
	var factory HeroicSkillFactory

	// Q技能
	factory = new(SteelTempestFactory)
	steelTempest := factory.CreateHeroicSkill()
	steelTempest.Desc()

	// W技能
	factory = new(WindWallFactory)
	windWall := factory.CreateHeroicSkill()
	windWall.Desc()

	// E技能
	factory = new(SweepingBladeFactory)
	sweepingBlade := factory.CreateHeroicSkill()
	sweepingBlade.Desc()

	// R技能
	factory = new(LastBreathFactory)
	lastBreath := factory.CreateHeroicSkill()
	lastBreath.Desc()

	// 这时候如果要添加一个技能或者重构一个技能只需要新增代码即可
	// TODO 被动技能
}
