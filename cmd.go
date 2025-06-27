package main

import (
	"os/exec"
	"strings"
)

func execCmd(name string, arg ...string) string {
	b, err := exec.Command(name, arg...).Output()
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(b))
}
