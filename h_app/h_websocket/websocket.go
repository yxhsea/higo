package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"github.com/gorilla/websocket"
	"net/http"
	"github.com/julienschmidt/httprouter"
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
			webSocketMap := viper.GetStringMap("websocket")
			readBuffer, _ := webSocketMap["read_buffer"].(int64)
			writeBuffer, _ := webSocketMap["write_buffer"].(int64)

			upgrader := websocket.Upgrader{
				ReadBufferSize:  int(readBuffer),
				WriteBufferSize: int(writeBuffer),
				CheckOrigin:     func(r *http.Request) bool { return true },
				Subprotocols:    []string{"binary"},
			}

			router := httprouter.New()
			router.GET("ws", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
				r.Header.Del("Sec-WebSocket-Protocol")
				ws, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println(ws)
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
