package logger

import (
	"flag"
	"io"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	flag.CommandLine.SetOutput(io.Discard)

	// set location of log file
	var logPath = "./logs/info.log"

	flag.Parse()
	var file, err1 = os.Create(logPath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logPath)
}
