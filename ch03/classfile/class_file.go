package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fileds       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) ConstantVersion() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fileds
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

//Parse()函数把[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

//read()方法依次调用其他方法解析class文件
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUnit16()
	self.thisClass = reader.readUnit16()
	self.superClass = reader.readUnit16()
	self.interfaces = reader.readUnit16s()
	self.fileds = readMembers(reader, self.constantPool)
	self.fileds = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

//ClassName()从常量池查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

//SuperClassName()从常量池查找超类名
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有Object没有超类
}

//InterfaceNames()从常量池查找接口名
func (self *ClassFile) InterfaceName() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

//0XCAFEBABE魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUnit32()
	if magic != 0XCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}

//检查版本号 JDK8
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUnit16()
	self.majorVersion = reader.readUnit16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 4950, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

