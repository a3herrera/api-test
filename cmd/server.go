package cmd

import (
	"context"
	"github.com/a3herrera/api-test/container"
	"github.com/a3herrera/api-test/container/logger"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"time"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: runServerCmd,
}

func runServerCmd(cmd *cobra.Command, args []string) {

	httpClient, err := container.NewServer()
	if err != nil {
		logger.Log.Errorf("Fail to retrieve all config: %s'", err)
	}

	go func() {
		if err := httpClient.ListAndServe(); err != nil {
			logger.Log.Errorf("Server error: %s", err)
		}
	}()
	logger.Log.Info("Server started")
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	<-stopChan // wait fo SIGINT
	logger.Log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(
		context.Background(),
		50*time.Millisecond)
	<-ctx.Done()
	cancel()
	logger.Log.Info("Server stopped")
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
