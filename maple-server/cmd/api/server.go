package api

import (
	"fmt"
	"log"
	"maple-server/pkg/logger"
	"maple-server/router"
	"maple-server/tools"
	config2 "maple-server/tools/config"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	aa := "1"
	logger.Infof("%s\n", aa)
}

func run() error {
	if mode != "" {
		config2.SetConfig(config, "settings.appliation.mode", mode)
	}

	if viper.GetString("settings.applition.mode") == string(tools.ModeProd) {

	}

	r := router.InitRouter()

	if port != "" {
		config2.SetConfig(config, "settings.appliation.port", port)
	}

	srv := &http.Server{
		Addr:    "",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Printf("%s Server Run http://%s:%s/ \r\n",
		tools.GetCurrntTimeStr(),
		config2.ApplicationConfig.Host,
		config2.ApplicationConfig.Port)

	return nil
}
