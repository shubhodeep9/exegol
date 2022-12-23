package api

import (
	"errors"
	"os"
	"os/exec"
)

func (ins *Installer) installPkg(uri, version string) (string, error) {
	getCmd := &exec.Cmd{
		Path:   ins.goExec,
		Args:   []string{ins.goExec, "install", uri + "@" + version},
		Env:    append(os.Environ(), "GO111MODULE=on", "GOBIN="+ins.goBin),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := getCmd.Run(); err != nil {
		return "", errors.New("error fetching package, ensure url is correct")
	}

	bins, err := os.ReadDir(ins.goBin)

	if err != nil || len(bins) == 0 {
		return "", errors.New("error finding package executable")
	}

	installedPath := ins.goBin + "/" + bins[0].Name()

	return installedPath, nil
}
