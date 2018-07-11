package classfile

import "encoding/binary"

type ClassReader struct {
	date [] byte
}

//读取u1型数据
func (self *ClassReader) readUnit8() uint8 {
	val := self.date[0]
	self.date = self.date[1:]
	return val
}

//读取u2型数据
func (self *ClassReader) readUnit16() uint16 {
	val := binary.BigEndian.Uint16(self.date)
	self.date = self.date[2:]
	return val
}

//读取u4型数据
func (self *ClassReader) readUnit32() uint32 {
	val := binary.BigEndian.Uint32(self.date)
	self.date = self.date[4:]
	return val
}

//读取u8型数据
func (self *ClassReader) readUnit64() uint64 {
	val := binary.BigEndian.Uint64(self.date)
	self.date = self.date[8:]
	return val
}

//读取u2型数据表
func (self *ClassReader) readUnit16s() []uint16 {
	n := self.readUnit16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUnit16()
	}
	return s
}

//读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.date[:n]
	self.date = self.date[n:]
	return bytes
}
