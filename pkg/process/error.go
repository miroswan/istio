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
	"os"
	"time"
)

// ExitError wraps an [os/exec.ExitError] and adds setters and getters
// for its struct members.
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
