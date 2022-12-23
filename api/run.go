package api

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (ins *Installer) RunPkg(installedPath, args string) {
	argSplit := append([]string{installedPath}, strings.Split(args, " ")...)
	getCmd := &exec.Cmd{
		Path:   installedPath,
		Args:   argSplit,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
	}

	if err := getCmd.Run(); err != nil {
		fmt.Println("Error", err)
	}
}
