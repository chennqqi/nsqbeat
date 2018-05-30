package main

import (
	"os"

	_ "output/nsq"
	_ "processor/command"

	"github.com/chennqqi/nsqbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
