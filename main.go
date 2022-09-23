package main

import (
	"os"
	"runtime"

	"github.com/JoseRodrigues443/is-my-team-awake/cmd"
	"github.com/JoseRodrigues443/is-my-team-awake/logger"
)

const (
	VERSION = "0.13"
)

func main() {
	logger.Log.Printf("Server v%s pid=%d started with processes: %d", VERSION, os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))
	cmd.Execute()
	logger.Log.Printf("Ending...")
}
