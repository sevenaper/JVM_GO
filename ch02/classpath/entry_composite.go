package classpath

import "strings"

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

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for_,entry:=range self{
		data,from,err := en
	}
}
