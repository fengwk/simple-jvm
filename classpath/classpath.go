package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	dirLib      = "lib"
	dirExt      = "ext"
	envJavaHome = "JAVA_HOME"
)

type Classpath struct {
	BootClasspath Entry
	ExtClasspath  Entry
	AppClasspath  Entry
}

// Parse 将jre路径和classpath解析为实际的Classpath
func Parse(jrePath, classpath string) (*Classpath, error) {
	jreAbsPath, err := getJreAbsPath(jrePath)
	if err != nil {
		return nil, err
	}

	bootClasspath, err := NewWildcardEntry(filepath.Join(jreAbsPath, dirLib))
	if err != nil {
		return nil, err
	}

	extClasspath, err := NewWildcardEntry(filepath.Join(jreAbsPath, dirLib, dirExt))
	if err != nil {
		return nil, err
	}

	appClasspath, err := parseClasspath(classpath)
	if err != nil {
		return nil, err
	}

	return &Classpath{BootClasspath: bootClasspath, ExtClasspath: extClasspath, AppClasspath: appClasspath}, nil
}

func getJreAbsPath(jrePath string) (string, error) {
	if jrePath == "" {
		jrePath = os.Getenv(envJavaHome)
	}

	jreAbsPath, err := filepath.Abs(jrePath)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(jreAbsPath); err != nil && os.IsNotExist(err) {
		return "", err
	}

	return jreAbsPath, nil
}

func parseClasspath(classpath string) (Entry, error) {
	var es []Entry
	for _, path := range strings.Split(classpath, string(os.PathListSeparator)) {
		if e, err := toEntry(path); err != nil {
			es = append(es, e)
		}
	}
	return NewCompositeEntry(es), nil
}

func toEntry(path string) (Entry, error) {
	if isZip(path) {
		return NewZipEntry(path)
	} else if isWildcard(path) {
		return NewWildcardEntry(path)
	} else {
		return NewDirEntry(path)
	}
}
