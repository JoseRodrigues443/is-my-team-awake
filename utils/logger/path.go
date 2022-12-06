package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

var (
	Log *log.Logger
)

// It returns the path of the directory containing the file that called it
func getPackagePath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}

// It returns the path to the project root directory
func GetProjectRootPath(logPath string) string {
	return fmt.Sprintf("%s/../%s", getPackagePath(), logPath)
}
