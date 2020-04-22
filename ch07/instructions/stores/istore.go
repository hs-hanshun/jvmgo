package stores

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type ISTORE struct {
	base.Index8Instruction
}

func (self *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, self.Index)
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
