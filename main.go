package main

import (
	"log"
	"os"

	"github.com/rmorris1218/conntest/cli"
)

func main() {
	app := cli.GetCli()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}