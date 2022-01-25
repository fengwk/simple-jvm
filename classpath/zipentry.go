package classpath

import (
	"archive/zip"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry 压缩文件条目
type ZipEntry struct {
	// 当前条目的绝对路径
	AbsZipPath string
}

// NewZipEntry 通过路径构建一个压缩文件条目，必须确保path指向的是一个zip文件
func NewZipEntry(path string) (*ZipEntry, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	rc, err := zip.OpenReader(absPath)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return &ZipEntry{AbsZipPath: absPath}, nil
}

// Read 通过classpath读取当前压缩条目中的文件字节流
func (ze *ZipEntry) Read(name string) ([]byte, error) {
	rc, err := zip.OpenReader(ze.AbsZipPath)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	for _, f := range rc.File {
		if f.Name == name {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			bytes, err := ioutil.ReadAll(rc)
			rc.Close()
			return bytes, err
		}
	}

	return nil, ErrClassNotFound
}
