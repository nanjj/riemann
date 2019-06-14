package main

import (
	"os/exec"
	"runtime"
	"time"
)

// open opens the specified URL in the default browser of the user.
func Open(url string) (err error) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	command := exec.Command(cmd, args...)
	if _, err = command.Output(); err == nil {
		time.Sleep(time.Second)
	}
	return
}
