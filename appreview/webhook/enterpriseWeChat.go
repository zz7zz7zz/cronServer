package webhook

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"open.com/cronServer/appreview/config"
	"open.com/cronServer/appreview/constant"
	"open.com/cronServer/appreview/models"
)

type EnterpriseWechat struct {
}

func (w EnterpriseWechat) OnWebHook(appReviewRecord *models.AppReviewRecord) {
	url := fmt.Sprintf(constant.AppStoreURL, appReviewRecord.Pkg)
	if appReviewRecord.Platform == constant.Android {
		url = fmt.Sprintf(constant.PlayStoreURL, appReviewRecord.Pkg)
	}

	tUTCCreate := time.Unix(int64(appReviewRecord.TaskCreateTs), 0).UTC()
	tUTCtUTCCreateStr := tUTCCreate.Format("2006年01月02日 15:04:05")

	tUTCApprove := time.Unix(int64(appReviewRecord.ApproveTs), 0).UTC()
	tUTCApproveStr := tUTCApprove.Format("2006年01月02日 15:04:05")

	channel := appReviewRecord.Channel
	if channel == "" {
		if appReviewRecord.Platform == constant.Android {
			channel = "gp"
		} else {
			channel = "appstore"
		}
	}

	content := fmt.Sprintf("---------------提醒：%s 审核通过---------------\n版本：%s\n包名：%s\n渠道：%s\n链接：%s\n任务创建时间(UTC)：%s\n审核通过时间(UTC)：%s\n\n审核开关已自动设置\n\n", appReviewRecord.Platform, appReviewRecord.Ver, appReviewRecord.Pkg, channel, url, tUTCtUTCCreateStr, tUTCApproveStr)
	resp, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+config.G_Config.Webhook.Wechat.Key, "application/json", strings.NewReader(`{"msgtype": "text", "text": {"content": "`+content+`","mentioned_list":["@all"]}}`))
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
