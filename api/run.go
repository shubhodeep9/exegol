package api

import (
	"fmt"
	"os"
	"os/exec"
)

func (ins *Installer) RunPkg(installedPath string) {
	getCmd := &exec.Cmd{
		Path:   installedPath,
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
	}

	if err := getCmd.Run(); err != nil {
		fmt.Println("Error", err)
	}
}
