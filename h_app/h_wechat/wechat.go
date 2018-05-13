package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/request"
	"gopkg.in/chanxuehong/wechat.v2/mp/menu"
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler core.Handler
	msgServer  *core.Server
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "Wechat",
		Short: "Wechat",
		Long:  "Wechat",
		Run: func(cmd *cobra.Command, args []string) {
			wechatMap := viper.GetStringMap("wechat")
			httpMap := viper.GetStringMap("http")

			host := httpMap["host"].(string)


			WxMpAppId := wechatMap["wx_mp_appid"].(string)
			WxMpAppSecret := wechatMap["wx_mp_appsecret"].(string)
			WxOriId := wechatMap["wx_ori_id"].(string)
			WxToken := wechatMap["wx_token"].(string)
			WxEncodedAesKey := wechatMap["wx_encoded_aes_key"].(string)
			WxMpAccessTokenServer := core.NewDefaultAccessTokenServer(WxMpAppId, WxMpAppSecret, nil)
			WxMpWechatClient := core.NewClient(WxMpAccessTokenServer, nil)
			fmt.Println(WxMpAppId, WxMpAppSecret, WxOriId, WxToken, WxEncodedAesKey, WxMpAccessTokenServer, WxMpWechatClient)

			router := httprouter.New()
			router.GET("/test",nil)

			mux := core.NewServeMux()
			mux.DefaultMsgHandleFunc(nil)
			mux.DefaultEventHandleFunc(nil)
			mux.MsgHandleFunc(request.MsgTypeText, nil)
			mux.EventHandleFunc(menu.EventTypeClick, nil)
			mux.EventHandleFunc(request.EventTypeScan, nil)
			mux.EventHandleFunc(request.EventTypeSubscribe, nil)

			msgHandler = mux
			msgServer = core.NewServer(WxOriId, WxMpAppId, WxToken, WxEncodedAesKey, msgHandler, nil)

			err := http.ListenAndServe(host, router)
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
		viper.SetConfigName("wechat")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./h_app/h_wechat")
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
