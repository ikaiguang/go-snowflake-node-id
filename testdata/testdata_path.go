package testdata

import (
	"path/filepath"
	"runtime"
)

// basePath is the root directory of this package.
var basePath string

const configPath = "./configs"

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basePath = filepath.Dir(currentFile)
}

// Path returns the absolute path the given relative file or directory path,
// relative to the google.golang.org/grpc/testdata directory in the user's GOPATH.
// If rel is already absolute, it is returned unmodified.
func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basePath, rel)
}

// ConfigPath 配置路径
func ConfigPath() string {
	return filepath.Join(basePath, configPath)
}
