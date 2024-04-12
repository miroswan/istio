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
	"context"
	"os/exec"
)

type Executor interface {
	Command(name string, args ...string) Cmd
	CommandContext(ctx context.Context, name string, arg ...string) Cmd
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

// NewExecutor returns an Executor.
func NewExecutor() Executor {
	return &executor{}
}
