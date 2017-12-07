package main

import (
	"os"

	"github.com/chennqqi/nsqbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
