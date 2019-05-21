package fileutil

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetExecPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func GetFilePath() string {
	// runtime.Caller()用户获取栈帧的信息, 参数为0表示当前函数, 参数为1表示上层调用该函数的函数.
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}
