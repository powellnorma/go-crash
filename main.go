package main

import (
	"os"

	"golang.org/x/sys/unix"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func redirectLogs(path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	f, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	chk(err)

	err = unix.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	chk(err)

	err = unix.Dup2(int(f.Fd()), int(os.Stdout.Fd()))
	chk(err)
}

func triggerCrash() {
	deps := make(map[int][]byte)
	for i := 1; i <= 1024*8; i++ {
		deps[i] = make([]byte, 1024)
	}
}

func main() {
	redirectLogs("/tmp/go-crash.log")
	triggerCrash()
}
