package classpath

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// DirEntry 文件夹条目
type DirEntry struct {
	AbsDirPath string
}

func NewDirEntry(path string) (*DirEntry, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, errors.Errorf("%s is not a dir", path)
	}

	return &DirEntry{AbsDirPath: absPath}, nil
}

func (de *DirEntry) Read(classpath string) ([]byte, error) {
	subPath := strings.ReplaceAll(classpath, string(NameSeparator), string(os.PathSeparator))
	fullPath := filepath.Join(de.AbsDirPath, subPath)
	return ioutil.ReadFile(fullPath)
}
