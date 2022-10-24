package cmd

import (
	"fmt"
	"maple-server/cmd/api"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "maple",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              "maple",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `欢迎使用 maple，可以使用 -h 查看命令`
		fmt.Printf("usageStr: %v\n", usageStr)
		// logger.Infof("%s\n", usageStr)
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
