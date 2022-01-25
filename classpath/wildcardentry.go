package classpath

import (
	"io/fs"
	"path/filepath"
)

// WildcardEntry 通配符条目将检测指定的文件夹目录下的所有jar包
type WildcardEntry struct {
	Entry *CompositeEntry
}

func NewWildcardEntry(path string) (*WildcardEntry, error) {
	baseDir, _ := filepath.Split(path)
	absBaseDir, err := filepath.Abs(baseDir)
	if err != nil {
		return nil, err
	}

	var es []Entry
	err = filepath.WalkDir(absBaseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && path != absBaseDir {
			// 跳过子目录
			return filepath.SkipDir
		}
		if isZip(path) {
			if ze, err := NewZipEntry(path); err == nil {
				es = append(es, ze)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &WildcardEntry{Entry: NewCompositeEntry(es)}, nil
}

func (we *WildcardEntry) Read(name string) ([]byte, error) {
	return we.Entry.Read(name)
}
