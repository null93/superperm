package main

import (
	"os"

	"github.com/null93/superperm/internal"
)

func main() {
	if err := internal.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
