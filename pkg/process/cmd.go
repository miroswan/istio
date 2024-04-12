// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package process

import (
	"io"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type Cmd interface {
	CmdBehaviors
	CmdGetters
	CmdSetters
}

type CmdGetters interface {
	Path() string
	Args() []string
	Env() []string
	Dir() string
	Stdin() io.Reader
	Stdout() io.Writer
	ExtraFiles() []*os.File
	SysProcAttr() *syscall.SysProcAttr
	ProcessState() *os.ProcessState
	Err() error
	Cancel() func() error
	WaitDelay() time.Duration
}

type CmdSetters interface {
	SetPath(path string)
	SetArgs(args []string)
	SetEnv(env []string)
	SetDir(dir string)
	SetStdin(stdin io.Reader)
	SetStdout(stdout io.Writer)
	SetExtraFiles(extraFiles []*os.File)
	SetSysProcAttr(sysProcAttr *syscall.SysProcAttr)
	SetProcessState(processState *os.ProcessState)
	SetErr(err error)
	SetCancel(cancel func() error)
	SetWaitDelay(waitDelay time.Duration)
}

type CmdBehaviors interface {
	CombinedOutput() ([]byte, error)
	Environ() []string
	Output() ([]byte, error)
	Run() error
	Start() error
	StderrPipe() (io.ReadCloser, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	String() string
	Wait() error
}

type cmd struct {
	*exec.Cmd
}

var _ Cmd = (*cmd)(nil)

func (c *cmd) Path() string                      { return c.Cmd.Path }
func (c *cmd) Args() []string                    { return c.Cmd.Args }
func (c *cmd) Env() []string                     { return c.Cmd.Env }
func (c *cmd) Dir() string                       { return c.Cmd.Dir }
func (c *cmd) Stdin() io.Reader                  { return c.Cmd.Stdin }
func (c *cmd) Stdout() io.Writer                 { return c.Cmd.Stdout }
func (c *cmd) ExtraFiles() []*os.File            { return c.Cmd.ExtraFiles }
func (c *cmd) SysProcAttr() *syscall.SysProcAttr { return c.Cmd.SysProcAttr }
func (c *cmd) ProcessState() *os.ProcessState    { return c.Cmd.ProcessState }
func (c *cmd) Err() error                        { return c.Cmd.Err }
func (c *cmd) Cancel() func() error              { return c.Cmd.Cancel }
func (c *cmd) WaitDelay() time.Duration          { return c.Cmd.WaitDelay }

func (c *cmd) SetPath(path string)                             { c.Cmd.Path = path }
func (c *cmd) SetArgs(args []string)                           { c.Cmd.Args = args }
func (c *cmd) SetEnv(env []string)                             { c.Cmd.Env = env }
func (c *cmd) SetDir(dir string)                               { c.Cmd.Dir = dir }
func (c *cmd) SetStdin(stdin io.Reader)                        { c.Cmd.Stdin = stdin }
func (c *cmd) SetStdout(stdout io.Writer)                      { c.Cmd.Stdout = stdout }
func (c *cmd) SetExtraFiles(extraFiles []*os.File)             { c.Cmd.ExtraFiles = extraFiles }
func (c *cmd) SetSysProcAttr(sysProcAttr *syscall.SysProcAttr) { c.Cmd.SysProcAttr = sysProcAttr }
func (c *cmd) SetProcessState(processState *os.ProcessState)   { c.Cmd.ProcessState = processState }
func (c *cmd) SetErr(err error)                                { c.Cmd.Err = err }
func (c *cmd) SetCancel(cancel func() error)                   { c.Cmd.Cancel = cancel }
func (c *cmd) SetWaitDelay(waitDelay time.Duration)            { c.Cmd.WaitDelay = waitDelay }

func (c *cmd) CombinedOutput() ([]byte, error)    { return c.Cmd.CombinedOutput() }
func (c *cmd) Environ() []string                  { return c.Cmd.Environ() }
func (c *cmd) Output() ([]byte, error)            { return c.Cmd.Output() }
func (c *cmd) Run() error                         { return c.Cmd.Run() }
func (c *cmd) Start() error                       { return c.Cmd.Start() }
func (c *cmd) StderrPipe() (io.ReadCloser, error) { return c.Cmd.StderrPipe() }
func (c *cmd) StdinPipe() (io.WriteCloser, error) { return c.Cmd.StdinPipe() }
func (c *cmd) StdoutPipe() (io.ReadCloser, error) { return c.Cmd.StdoutPipe() }
func (c *cmd) String() string                     { return c.Cmd.String() }
func (c *cmd) Wait() error                        { return c.Cmd.Wait() }

func ConvertCmd(execCmd *exec.Cmd) Cmd {
	return &cmd{execCmd}
}
