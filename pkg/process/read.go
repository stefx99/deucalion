package process

import (
	"sync"

	"github.com/prometheus/procfs"
	"github.com/stefx99/deucalion/internal/logger"
)

var procFS procfs.FS
var onceFS sync.Once

func init() {
	onceFS.Do(func() {
		var err error
		procFS, err = procfs.NewDefaultFS()
		if err != nil {
			logger.Error("Unable to read default proc mount point")
		}
	})
}

func getAllProcesses() []procfs.Proc {
	procs, err := procFS.AllProcs()
	if err != nil {
		logger.Fatal("Unable to fetch processes")
	}

	return procs
}

func findProcessByCommand(command string) procfs.Proc {
	for _, proc := range getAllProcesses() {
		cmdLine, _ := proc.CmdLine()
		if command == cmdLine[0] {
			return proc
		}
	}
	return procfs.Proc{} // Return zero value Proc if no matching process found
}
