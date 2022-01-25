package classpath

import "errors"

// ErrClassNotFound 查找不到类相应的异常
var ErrClassNotFound = errors.New("class not found")

// NameSeparator name路径所使用的分割符
const NameSeparator = '/'

// Entry 条目，是资源的抽象接口，可以使用classpath访问Entry中的资源
type Entry interface {

	// Read 通过name从当前条目中读取字节流
	// name使用NameSeparator分割
	// 在找不到相应name对应资源时返回ErrClassNotFound
	Read(name string) ([]byte, error)
}
