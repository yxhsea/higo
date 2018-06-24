package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
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

			//路由解析
			router := gin.Default()
			router.GET("/test", func(context *gin.Context) {
				fmt.Println("The is hihello testing...")
			})
			err := router.Run(host)
			if err != nil {
				fmt.Println(err)
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
