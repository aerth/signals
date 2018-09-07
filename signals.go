// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package signals provides a string-to-signal map
package signals

import (
	"strings"
	"syscall"
)

func Signal(s string) syscall.Signal {
	s = strings.ToUpper(s)
	s = strings.TrimPrefix(s, "SIG")
	s = "SIG" + s
	return Signals[s]
}

var Signals = map[string]syscall.Signal{
	"SIGABRT":   syscall.Signal(0x6),
	"SIGALRM":   syscall.Signal(0xe),
	"SIGBUS":    syscall.Signal(0x7),
	"SIGCHLD":   syscall.Signal(0x11),
	"SIGCLD":    syscall.Signal(0x11),
	"SIGCONT":   syscall.Signal(0x12),
	"SIGFPE":    syscall.Signal(0x8),
	"SIGHUP":    syscall.Signal(0x1),
	"SIGILL":    syscall.Signal(0x4),
	"SIGINT":    syscall.Signal(0x2),
	"SIGIO":     syscall.Signal(0x1d),
	"SIGIOT":    syscall.Signal(0x6),
	"SIGKILL":   syscall.Signal(0x9),
	"SIGPIPE":   syscall.Signal(0xd),
	"SIGPOLL":   syscall.Signal(0x1d),
	"SIGPROF":   syscall.Signal(0x1b),
	"SIGPWR":    syscall.Signal(0x1e),
	"SIGQUIT":   syscall.Signal(0x3),
	"SIGSEGV":   syscall.Signal(0xb),
	"SIGSTKFLT": syscall.Signal(0x10),
	"SIGSTOP":   syscall.Signal(0x13),
	"SIGSYS":    syscall.Signal(0x1f),
	"SIGTERM":   syscall.Signal(0xf),
	"SIGTRAP":   syscall.Signal(0x5),
	"SIGTSTP":   syscall.Signal(0x14),
	"SIGTTIN":   syscall.Signal(0x15),
	"SIGTTOU":   syscall.Signal(0x16),
	"SIGUNUSED": syscall.Signal(0x1f),
	"SIGURG":    syscall.Signal(0x17),
	"SIGUSR1":   syscall.Signal(0xa),
	"SIGUSR2":   syscall.Signal(0xc),
	"SIGVTALRM": syscall.Signal(0x1a),
	"SIGWINCH":  syscall.Signal(0x1c),
	"SIGXCPU":   syscall.Signal(0x18),
	"SIGXFSZ":   syscall.Signal(0x19),
}
