package webhook

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type EnterpriseWechat struct {
}

func (w EnterpriseWechat) OnWebHook() {
	key := "替换为你的企业微信机器人 Key"
	content := fmt.Sprintf("平台：%s\n版本：%s\n包名：%s\n渠道：%s\n结果：审核通过", "android", "1.0.0", "com.inhobchat.hobicat", "GooglePlay")
	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key, "application/json", strings.NewReader(`{"msgtype": "text", "text": {"content": "`+content+`","mentioned_list":["@all"]}}`))
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 打印响应内容
	fmt.Println("响应内容:", string(body))
}
