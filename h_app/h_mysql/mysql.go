package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"github.com/gohouse/gorose"
	_ "github.com/go-sql-driver/mysql"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "MysqlServer",
		Short: "MysqlServer",
		Long:  "MysqlServer",
		Run: func(cmd *cobra.Command, args []string) {
			//加载数据库配置文件信息
			databaseMap := viper.GetStringMap("database")
			host, _ := databaseMap["host"].(string)
			port, _ := databaseMap["port"].(string)
			user, _ := databaseMap["user"].(string)
			password, _ := databaseMap["password"].(string)
			dbname, _ := databaseMap["dbname"].(string)
			charset, _ := databaseMap["charset"].(string)
			poolnum, _ := databaseMap["pollnum"].(int64)

			//Db配置
			var DbConfig = map[string]interface{}{
				"Default": "mysql_dev",
				"SetMaxOpenConns": poolnum,
				"SetMaxIdleConns": 10,
				"Connections": map[string]map[string]string{
					"mysql_dev": map[string]string{
						"host":     host,
						"username": user,
						"password": password,
						"port":     port,
						"database": dbname,
						"charset":  charset,
						"protocol": "tcp",
						"prefix":   "",
						"driver":   "mysql",
					},
				},
			}

			//dbr连接资源句柄
			dbr, err := gorose.Open(DbConfig)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(dbr)
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
		viper.SetConfigName("mysql")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./h_app/h_mysql")
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
