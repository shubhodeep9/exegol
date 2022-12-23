package api

import (
	"errors"
	"os"
	"os/exec"
)

type Installer struct {
	goBin  string
	goExec string
}

func InitApi(uri, version, args string) error {
	path, err := exec.LookPath("go")

	if err != nil {
		return errors.New("go executable not locatable, please ensure go install and GOPATH is set")
	}

	installer := Installer{
		goBin:  os.TempDir() + "/exegol/bin",
		goExec: path,
	}

	installer.PreCleanup()

	installedPath, err := installer.installPkg(uri, version)

	if err != nil {
		return err
	}

	installer.ExitCleanup(installedPath)

	installer.RunPkg(installedPath, args)

	installer.PostCleanup(installedPath)

	return nil
}
