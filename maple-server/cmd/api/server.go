package api

import (
	"fmt"
	"log"
	config2 "maple-server/tools/config"

	"github.com/spf13/cobra"
)

var (
	config   string
	port     string
	mode     string
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "maple server config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func usage() {
	usageStr := `starting api server`
	log.Printf("%s\n", usageStr)
}

func setup() {
	fmt.Printf("config: %v\n", config)
	// 1. 读取配置
	config2.ConfigSetup(config)

}

func run() error {
	fmt.Printf("2222")
	return nil
}
