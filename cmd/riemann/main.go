package main

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use: "riemann",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		RootCmd.Println(err)
	}
}
