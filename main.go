package main

import (
	"os"

	"github.com/linuxsuren/atest-ext-store-database/cmd"
)

func main() {
	cmd := cmd.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
