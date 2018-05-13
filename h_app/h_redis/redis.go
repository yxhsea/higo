package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
	"github.com/garyburd/redigo/redis"
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
			redisMap := viper.GetStringMap("redis")
			host, _ := redisMap["host"].(string)
			auth, _ := redisMap["auth"].(string)
			poolNum, _ := redisMap["poolnum"].(int)

			//redisConn资源连接句柄
			redisConn := &redis.Pool{
				MaxIdle:     poolNum,
				MaxActive:   poolNum,
				IdleTimeout: 240 * time.Second,
				Dial: func() (redis.Conn, error) {
					c, err := redis.Dial("tcp", host)
					if err != nil {
						return nil, err
					}
					if auth != "" {
						if _, err := c.Do("AUTH", auth); err != nil {
							c.Close()
							return nil, err
						}
					}
					return c, err
				},
				TestOnBorrow: func(c redis.Conn, t time.Time) error {
					_, err := c.Do("PING")
					return err
				},
			}

			//Example
			conn := redisConn.Get()
			defer conn.Close()
			valueInt, err := redis.Int(conn.Do("SET","key","value"))
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(valueInt)
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
