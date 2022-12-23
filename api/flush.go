package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func (ins *Installer) ExitCleanup(path string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("attempting to remove binary")
		os.Remove(path)
		os.Exit(0)
	}()
}

func (ins *Installer) PreCleanup() {
	// cleaning up goBinDir
	if err := os.RemoveAll(ins.goBin); err != nil {
		fmt.Println("Error cleaning up bin dir", err)
		return
	}
}

func (ins *Installer) PostCleanup(path string) {
	os.Remove(path)
}
