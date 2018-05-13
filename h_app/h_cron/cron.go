package main

import (
	"os"
	"fmt"
	"time"
	"path/filepath"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/fsnotify/fsnotify"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "CoinCron",
		Short: "CoinCron of CoinCloud",
		Long:  "CoinCron is cron part of CoinCloud",
		Run: func(cmd *cobra.Command, args []string) {
			//定时任务
			c := cron.New()
			//每日凌晨，0点1分
			c.AddFunc("0 1 0 * * *", func() {
				fmt.Println("~~~~")
			})
			c.AddFunc("*/5 * * * * *", func() {
				fmt.Println("~~~~")
			})

			c.Start()
			defer c.Stop()

			t1 := time.NewTimer(time.Second * 10)
			for {
				select {
				case <-t1.C:
					t1.Reset(time.Second * 10)
				}
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
	}else{
		viper.SetConfigName("cron")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./h_app/h_cron")
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
