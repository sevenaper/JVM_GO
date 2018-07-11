package classfile
/*
常量池实际上也是一个表，但是有三点需要特别注意。第一，表头给出的常量池
大小比实际大1.假设表头给出的值为n，那么常量池实际的大小是n-1.第二，有效的
常量池索引是1~n-1，0是无效索引，表示不指向任何常量。第三，CONSTANT_Long_info
和CONSTANT_Double_info各占两个位置。也就是说，如果常量池中存在这两种常量，
实际的常量数量比n-1还要少，而且1~n-1某些数会变成无效索引。
 */
type ConstantPool []ConstantInfo
//读取常量池
func readConstantPool(reader*ClassReader)ConstantPool  {
	cpCount := int(reader.readUnit16())
	cp := make([]ConstantInfo,cpCount)
	for i:= 1;i<cpCount;i++{//索引从1开始
		cp[i] = readConstantInfo(reader,cp)
		switch cp[i].(type) {
		case*ConstantLongInfo,*ConstantDoubleInfo:
			i++//各占据两个位置
		}
	}
	return cp
}
//按照索引查找常量
func(self ConstantPool)getConstantInfo(index uint16)ConstantInfo{
	if cpInfo:= self[index];cpInfo!=nil{
		return cpInfo
	}
	panic("Invalid constant pool index!")
}
//从常量池查找字段和方法的名字和描述符
func (self ConstantPool)getNameAndType (index uint16)(string,string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name,_type
}
//从常量池找类名
func (self ConstantPool)getClassName(index uint16)string{
	classInfo:=self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

//从常量池中查找UTF-8字符串
func (self ConstantPool)getUtf8(index uint16)string  {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}