package classpath

import (
	"strings"
	"errors"
)

type CompositeEntry [] Entry

//CompositeEntry由更小的Entry组成，按分割符把参数按分割符分成小路径，然后
//每个小路径都转化成Entry实例
func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

//依次调用每一个子路径的readClass()方法
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

//调用每一个子路径的String()方法，然后把得到的字符串用路径分割副拼接起来
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
