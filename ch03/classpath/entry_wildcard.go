package classpath

import (
	"os"
	"path/filepath"
	"strings"
)
//首先把路径末尾的*去掉，得到baseDir然后调用filepath包
//的Walk()函数遍历baseDir创建ZipEntry。
func newWildcardEntry(path string)CompositeEntry  {
	baseDir := path[:len(path)-1]//remove *
	var compositeEntry []Entry
	walkFn := func(path string,info os.FileInfo,err error)error{
		if err !=nil{
			return err
		}
		if info.IsDir() && path != baseDir{
			return filepath.SkipDir
		}
		if strings.HasSuffix(path,".jar")||strings.HasSuffix(path,".JAR"){
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry,jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir,walkFn)
	return compositeEntry
}