package msg

import (
	"fmt"
	"os"
	"path"
	"time"
)

var (
	VerboseLogging bool // if set allow verbose logging to be shown
	DebugLogging   bool // if set allow debug logging to be shown
	Testing        bool // if set disable logging
)

// echo "$(date +"%Y-%m-%d %H:%M:%S") $(hostname) $(basename $0)[$$]: $@"
// FIXME: this is all inefficient, unclean and needs fixing later especially as done on each call.
func Info(format string, a ...interface{}) {
	myhostname, _ := os.Hostname()
	_, myname := path.Split(os.Args[0])

	fmt.Printf(
		"%v %v %v[%v]: %v\n",
		time.Now().Format("2006-01-02 15:04:05"),
		myhostname,
		myname,
		os.Getpid(),
		fmt.Sprintf(format, a...),
	)
}

// Verbose only writes messages if VerboseLogging = true
func Verbose(format string, a ...interface{}) {
	if !VerboseLogging {
		return
	}

	Info("VERBOSE "+format, a...)
}

// Warning writes a warning message
func Warning(format string, a ...interface{}) {
	Info("WARNING "+format, a...)
}

// Exit writes a message and exits with the given exit code.
func Exit(rc int, format string, a ...interface{}) {
	Info("ERROR "+format, a...)
	os.Exit(rc)
}

// Debug only writes messages if DebugLogging = true
func Debug(format string, a ...interface{}) {
	if !DebugLogging {
		return
	}

	Info("DEBUG "+format, a...)
}
