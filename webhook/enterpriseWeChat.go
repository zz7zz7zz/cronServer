package webhook

import (
	"net/http"
	"strings"
)

func SendTextMessage(key string, content string) {
	http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key, "application/json", strings.NewReader(`{"msgtype": "text", "text": {"content": "`+content+`","mentioned_list":["@all"]}}`))
}
