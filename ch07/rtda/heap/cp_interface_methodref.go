package heap

import "jvmgo/ch07/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(pool *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.pool = pool
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}