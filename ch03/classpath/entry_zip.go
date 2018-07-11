package classpath

import (
	"path/filepath"
	"archive/zip"
	"errors"
	"io/ioutil"
)

type ZipEntry struct {
	absPath string
}

//先把参数转化成绝对路径，如果转换过程中出现错误，则调用
//panic()函数种植程序执行，否则创建DirEntry实例并返回
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

//返回变量的字符串表示
func (self *ZipEntry) String() string {
	return self.absPath
}

//从ZIP文件中提取class文件,首先打开ZIP文件，若这部出错，直接返回。
//然后遍历ZIP压缩包里面的文件，看是否能找到class文件。如果可以找到，读出来并返回。
//如果找不到，或者出现其他错误，那么返回错误信息。可以进行优化
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
