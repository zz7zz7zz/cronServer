package webhook

import (
	"cronServer/config"
	"cronServer/constant"
	"cronServer/models"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type EnterpriseWechat struct {
}

func (w EnterpriseWechat) OnWebHook(appReviewRecord *models.AppReviewRecord) {
	url := fmt.Sprintf(constant.AppStoreURL, appReviewRecord.Pkg)
	if appReviewRecord.Platform == constant.Android {
		url = fmt.Sprintf(constant.PlayStoreURL, appReviewRecord.Pkg)
	}
	content := fmt.Sprintf("平台：%s\n版本：%s\n包名：%s\n渠道：\n%s链接：%s\n结果：审核通过", appReviewRecord.Platform, appReviewRecord.Ver, appReviewRecord.Pkg, appReviewRecord.Channel, url)
	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+config.GConfig.Webhook.Wechat.Key, "application/json", strings.NewReader(`{"msgtype": "text", "text": {"content": "`+content+`","mentioned_list":["@all"]}}`))
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
	fmt.Println("EnterpriseWechat 响应内容:", string(body))
}
