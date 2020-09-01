package main

import (
	"fmt"
	"os"
	"os/exec"
)

func installPkg (uri string) {
	path, err := exec.LookPath("go")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	getCmd := &exec.Cmd{
		Path: path,
		Args: []string{ path, "get", "-u", uri},
		Env: append(os.Environ(), "GO111MODULE=on"),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := getCmd.Run(); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	path, err = exec.LookPath("lazyme")

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	getCmd = &exec.Cmd{
		Path: path,
		Args: []string{ path },
		Stdout: os.Stdout,
		Stdin: os.Stdin,
		Stderr: os.Stderr,
	}

	if err := getCmd.Run(); err != nil {
		fmt.Println("Error", err)
	}

}

