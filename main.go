package main

import (
	"os"

	"github.com/sikender/xkcd/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
