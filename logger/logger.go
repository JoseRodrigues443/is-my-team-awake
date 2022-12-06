package logger

import (
	"flag"
	"io"
	"log"
	"os"

	utils "github.com/JoseRodrigues443/is-my-team-awake/utils/logger"
)

var (
	Log *log.Logger
)

func init() {
	flag.CommandLine.SetOutput(io.Discard)

	// set location of log file
	// @JoseRodrigues443 TODO add these dir to config file
	var logPath = "./logs/info.log"

	var file, err = os.Create(utils.GetProjectRootPath(logPath))

	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + file.Name())
}
