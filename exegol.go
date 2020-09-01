package main

import (
	"fmt"
	"os"
)

func exegol(pkg string) {
	installPkg(pkg)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("arguments missing bro")
		return
	}

	var (
		pkg = os.Args[1]
	)

	exegol(pkg)
}
