package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"github.com/eclipse/paho.mqtt.golang"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "MqttServer",
		Short: "MqttServer",
		Long:  "MqttServer",
		Run: func(cmd *cobra.Command, args []string) {
			//加载配置文件信息
			mqttMap := viper.GetStringMap("mqtt")
			broker, _ := mqttMap["host"].(string)
			user, _ := mqttMap["user"].(string)
			password, _ := mqttMap["password"].(string)
			clientId, _ := mqttMap["clientId"].(string)
			store, _ := mqttMap["store"].(string)

			//配置Mqtt
			opts := mqtt.NewClientOptions()
			opts.AddBroker(broker)
			opts.SetClientID(clientId)
			opts.SetUsername(user)
			opts.SetPassword(password)
			opts.SetCleanSession(false)
			if store != ":memory:" {
				opts.SetStore(mqtt.NewFileStore(store))
			}

			//实例化一个Mqtt客户端
			client := mqtt.NewClient(opts)
			if token := client.Connect(); token.Wait() && token.Error() != nil {
				fmt.Println(token.Error())
			}

			//topic解析
			client.AddRoute("topic", func(client mqtt.Client, message mqtt.Message) {
				//接收到的数据包
				fmt.Println(message.Payload())
			})
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
		viper.SetConfigName("mqtt")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./h_app/h_mqtt")
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
