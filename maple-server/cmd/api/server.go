package api

import (
	"context"
	"fmt"
	"log"
	"maple-server/database"
	"maple-server/global/orm"
	"maple-server/pkg/logger"
	"maple-server/router"
	"maple-server/tools"
	config2 "maple-server/tools/config"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
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
	// 2. 初始化数据库链接
	database.Setup()
}

func run() error {
	if mode != "" {
		config2.SetConfig(config, "settings.application.mode", mode)
	}

	if viper.GetString("settings.application.mode") == string(tools.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()

	defer func() {
		err := orm.Eloquent.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	if port != "" {
		config2.SetConfig(config, "settings.application.port", port)
	}

	srv := &http.Server{
		Addr:    config2.ApplicationConfig.Host + ":" + config2.ApplicationConfig.Port,
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
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", tools.GetCurrntTimeStr())

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", tools.GetCurrntTimeStr())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")

	return nil
}
