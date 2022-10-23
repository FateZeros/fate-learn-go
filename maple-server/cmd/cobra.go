package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:	"maple",
	Run: 	func(cmd *cobra.Command, args []string) {
				usageStr := `欢迎`
				logger.Infof("%s\n", usageStr)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}