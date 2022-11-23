package logger

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
func getProjectRootPath(logPath string) string {
	return fmt.Sprintf("%s/../%s", getPackagePath(), logPath)
}

func init() {
	flag.CommandLine.SetOutput(io.Discard)

	// set location of log file
	// @JoseRodrigues443 TODO add these dir to config file
	var logPath = "./logs/info.log"

	var file, err = os.Create(getProjectRootPath(logPath))

	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + file.Name())
}
