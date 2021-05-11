package cmd

import (
	"fmt"
	"github.com/a3herrera/api-test/container/service"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "searcher",
	Short: "Search centralizer",
}

func Execute() {
	_ = service.InitApp()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println()
		os.Exit(1)
	}
}
