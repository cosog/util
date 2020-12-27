// dir
package util_path

import (
	"os"
	"runtime"
)

var ostype = runtime.GOOS

func GetProjectPath() string {
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}
func GetCurrentPath() string {
	path := GetProjectPath()
	if ostype == "windows" {
		path = path + "\\"
	} else if ostype == "linux" {
		path = path + "/"
	}
	return path
}
func GetConfigPath() string {
	path := GetProjectPath()
	if ostype == "windows" {
		path = path + "\\" + "config\\"
	} else if ostype == "linux" {
		path = path + "/" + "config/"
	}
	return path
}

func GetConLogPath() string {
	path := GetProjectPath()
	if ostype == "windows" {
		path = path + "\\log\\"
	} else if ostype == "linux" {
		path = path + "/log/"
	}
	return path
}
