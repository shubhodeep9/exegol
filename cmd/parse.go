package cmd

import (
	"flag"
	"log"

	"github.com/shubhodeep9/exegol/api"
)

func Parse() {
	uri := flag.String("uri", "", "provide url of package to be installed")
	version := flag.String("v", "latest", "provide version of the package to be installed, i.e latest")
	args := flag.String("arg", "", "provide arguments to be provided to package")

	flag.Parse()

	if err := api.InitApi(*uri, *version, *args); err != nil {
		log.Fatal(err)
	}
}
