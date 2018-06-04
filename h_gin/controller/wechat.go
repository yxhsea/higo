package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/menu"
	"log"
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler core.Handler
	msgServer  *core.Server
)

func SetWxRouter() {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)

	msgHandler = mux
	msgServer = core.NewServer("", "wx592b2759cf52d2ca", "yxhsea", "", msgHandler, nil)
}

func WxCallbackHandler(c *gin.Context) {
	msgServer.ServeHTTP(c.Writer, c.Request, nil)
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func CreateMenu(ctx *gin.Context) {
	WxMpAccessTokenServer := core.NewDefaultAccessTokenServer("wx592b2759cf52d2ca", "100eb8ca944a10058415fc2c069fdd0d", nil)
	WxMpWechatClient := core.NewClient(WxMpAccessTokenServer, nil)

	MenuList := &menu.Menu{
		Buttons: []menu.Button{
			menu.Button{
				Type: "view", // 非必须; 菜单的响应动作类型
				//Name: "设备绑定",                               // 必须;  菜单标题
				Name: "测试页面",                               // 必须;  菜单标题
				Key:  "",                                   // 非必须; 菜单KEY值, 用于消息接口推送
				URL:  "http://1618512a3z.iask.in:39364/h5", // 非必须; 网页链接, 用户点击菜单可打开链接
				//URL:        "http://wawadevicebind.tunnel.aioil.cn/dist/#/device_list", // 非必须; 网页链接, 用户点击菜单可打开链接
				MediaId:    "",              // 非必须; 调用新增永久素材接口返回的合法media_id
				AppId:      "",              // 非必须; 跳转到小程序的appid
				PagePath:   "",              // 非必须; 跳转到小程序的path
				SubButtons: []menu.Button{}, // 非必须; 二级菜单数组
			},
			menu.Button{
				Type:       "view",                                                                                                  // 非必须; 菜单的响应动作类型
				Name:       "生产",                                                                                                    // 必须;  菜单标题
				Key:        "",                                                                                                      // 非必须; 菜单KEY值, 用于消息接口推送
				URL:        "https://wawafront.tunnel.aioil.cn/dist/?from=singlemessage&isappinstalled=0#/?merchant_id=20000332793", // 非必须; 网页链接, 用户点击菜单可打开链接
				MediaId:    "",                                                                                                      // 非必须; 调用新增永久素材接口返回的合法media_id
				AppId:      "",                                                                                                      // 非必须; 跳转到小程序的appid
				PagePath:   "",                                                                                                      // 非必须; 跳转到小程序的path
				SubButtons: []menu.Button{},                                                                                         // 非必须; 二级菜单数组
			},
			menu.Button{
				Type:       "view",                                                                                                 // 非必须; 菜单的响应动作类型
				Name:       "开发",                                                                                                   // 必须;  菜单标题
				Key:        "",                                                                                                     // 非必须; 菜单KEY值, 用于消息接口推送
				URL:        "https://wawafront.tunnel.aioil.cn/dev/?from=singlemessage&isappinstalled=0#/?merchant_id=20000332793", // 非必须; 网页链接, 用户点击菜单可打开链接
				MediaId:    "",                                                                                                     // 非必须; 调用新增永久素材接口返回的合法media_id
				AppId:      "",                                                                                                     // 非必须; 跳转到小程序的appid
				PagePath:   "",                                                                                                     // 非必须; 跳转到小程序的path
				SubButtons: []menu.Button{},                                                                                        // 非必须; 二级菜单数组
			},
		},
	}
	err := menu.Create(WxMpWechatClient, MenuList)
	if err != nil {
		fmt.Println("______", err.Error())
	}
	fmt.Println("==========")
}
