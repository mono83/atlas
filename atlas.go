package main

import (
	"fmt"
	"github.com/mono83/atlas/cmd"
	"os"
)

func main() {
	if err := cmd.AtlasCmd.Execute(); err != nil {
		fmt.Println("Execution error occured:", err)
		os.Exit(1)
	}
}
