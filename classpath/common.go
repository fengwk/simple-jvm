package classpath

import "path/filepath"

const (
	extJar   = ".jar"
	extWar   = ".war"
	extZip   = ".zip"
	wildcard = "*"
)

func isZip(path string) bool {
	ext := filepath.Ext(path)
	return ext == extJar || ext == extWar || ext == extZip
}

func isWildcard(path string) bool {
	_, f := filepath.Split(path)
	return f == wildcard
}
