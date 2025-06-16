//go:build !windows

package tasks

import "syscall"

func getSysProcAttr() *syscall.SysProcAttr {
	return nil
}
