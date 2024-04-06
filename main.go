package main

import (
	"os"

	"github.com/DarkZoneSD/vmSetup/src/args"
)

func main() {
	args.HandleArgs(os.Args)
}
