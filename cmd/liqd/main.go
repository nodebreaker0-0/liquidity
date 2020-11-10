package main

import (
	"os"

	"github.com/nodebreaker0-0/liquidity/cmd/liqd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
