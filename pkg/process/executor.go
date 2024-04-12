package process

import (
	"context"
	"os/exec"
)

type Executor interface {
	Command(name string, args ...string) Cmd
	CommandContext(ctx context.Context, name string, arg ...string) Cmd
	LookPath(file string) (string, error)
}

type executor struct{}

var _ Executor = (*executor)(nil)

// Command returns a Cmd for execution.
func (e *executor) Command(name string, args ...string) Cmd {
	return ConvertCmd(exec.Command(name, args...))
}

// CommandContext returns a Cmd for execution with a given context.
func (e *executor) CommandContext(ctx context.Context, name string, arg ...string) Cmd {
	return ConvertCmd(exec.CommandContext(ctx, name, arg...))
}

func (e *executor) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

// NewExecutor returns an Executor.
func NewExecutor() Executor {
	return &executor{}
}
