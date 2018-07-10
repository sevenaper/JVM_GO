package classpath

import (
	"path/filepath"
	"io/ioutil"
)

type DirEntry struct {
	absDir string
}

//先把参数转化成绝对路径，如果转换过程中出现错误，则调用
//panic()函数种植程序执行，否则创建DirEntry实例并返回
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

//readClass()先把目录和class文件名拼成一个完整的路径
//然后调用ioutil包提供的ReadFile()函数读取class文件
//内容，最后返回。
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

//返回变量的字符串表示
func (self *DirEntry) String() string {
	return self.absDir

}
