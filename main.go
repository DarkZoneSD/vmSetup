package main

import (
	"os"
	"vmSetup/src/args"
)

func main() {
	args.HandleArgs(os.Args)
}
