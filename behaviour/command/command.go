package main

import "fmt"

// Doctor 命令接受者
type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// Command 抽象命令
type Command interface {
	Treat()
}

// CommandTreatEye 治疗眼睛的病单-具体命令的执行
type CommandTreatEye struct {
	doctor *Doctor
}

func (cmd *CommandTreatEye) Treat() {
	cmd.doctor.treatEye()
}

// CommandTreatNose 治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// Nurse 护士-调用命令者
type Nurse struct {
	CmdList []Command // 收集命令的集合
}

func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		//执行病单绑定的命令(这里会调用病单已经绑定的医生的诊断方法)
		cmd.Treat()
	}
}

func main() {
	// 命令接收者 - 依赖病单，通过填写病单，让医生看病
	docker := new(Doctor)
	// 治疗眼睛的病单
	cmdEye := CommandTreatEye{doctor: docker}
	// 治疗鼻子的病单
	cmdNose := CommandTreatNose{doctor: docker}

	// 护士-命令调用者
	nurse := new(Nurse)
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)

	// 执行病单
	nurse.Notify()
}
