package process

import (
	"os"
	"time"
)

type ExitError interface {
	error
	ExitCode() int
	Exited() bool
	Pid() int
	String() string
	Success() bool
	Sys() any
	SysUsage() any
	SystemTime() time.Duration
	UserTime() time.Duration

	// Properties
	ExitErrorGetters
	ExitErrorSetters
}

type ExitErrorGetters interface {
	ProcessState() *os.ProcessState
	Stderr() []byte
}

type ExitErrorSetters interface {
	SetProcessState(processState *os.ProcessState)
	SetStderr(stderr []byte)
}
