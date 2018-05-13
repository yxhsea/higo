package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "HttpServer",
		Short: "HttpServer",
		Long:  "HttpServer",
		Run: func(cmd *cobra.Command, args []string) {
			//加载配置文件信息
			httpMap := viper.GetStringMap("http")
			host, _ := httpMap["host"].(string)
			timeout, _ := httpMap["timeout"].(time.Duration)

			//路由解析
			router := fasthttprouter.New()
			router.GET("/test", middleware(func(ctx *fasthttp.RequestCtx) {
				fmt.Println("hello")
				return
			}))

			//服务配置
			server := &fasthttp.Server{
				Name:         "HttpServer",
				ReadTimeout:  timeout * time.Second,
				WriteTimeout: timeout * time.Second,
				Handler:      router.Handler,
			}

			//服务监听
			err := server.ListenAndServe(host)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}

func initConfig() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("http")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./h_app/h_http")
		viper.AddConfigPath(dir)
		viper.AutomaticEnv()
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

func middleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		fmt.Println("middleware")
		return
	})
}