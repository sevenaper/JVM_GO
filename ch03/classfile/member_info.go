package classfile

type MemberInfo struct {
	cp              ConstantPool    //保存常量指针
	accessFLags     uint16          //访问标志
	nameIndex       uint16          //字段名或者方法名
	descriptorIndex uint16          //字段或者方法的描述符
	attributes      []AttributeInfo //属性表
}

//读取字段表或者方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUnit16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMembers(reader, cp)
	}
	return members
}

//读取字段或方法区
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFLags:     reader.readUnit16(),
		nameIndex:       reader.readUnit16(),
		descriptorIndex: reader.readUnit16(),
		attributes:      readAttributes(reader, cp),
	}
}

//从常量池查找字段或者方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

//从常量池查找字段或者方法描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
