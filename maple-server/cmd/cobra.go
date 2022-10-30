package cmd

import (
	"errors"
	"fmt"
	"maple-server/cmd/api"
	"maple-server/pkg/logger"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "maple",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              "maple long",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `欢迎使用 maple，可以使用 -h 查看命令`
		fmt.Printf("usageStr: %v\n", usageStr)
		logger.Infof("%s\n", usageStr)
	},
}

func init() {
	fmt.Printf("cobra init \n")
	rootCmd.AddCommand(api.StartCmd)
	// rootCmd.AddCommand(migrate.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
