package signals

import (
	"syscall"
	"testing"
)

func TestString2Signal(t *testing.T) {
	type testcase struct {
		try  string
		gold syscall.Signal
	}
	tc := []testcase{
		{"HUP", syscall.SIGHUP},
		{"SIGHUP", syscall.SIGHUP},
		{"USR1", syscall.SIGUSR1},
		{"SIGUSR1", syscall.SIGUSR1},
	}
	for i := range tc {
		if sig := Signal(tc[i].try); sig != tc[i].gold {
			t.Logf("wanted %v, tried %s, got %v", sig, tc[i].try, tc[i].gold)
			t.Fail()
		}
	}
}
