protoc --go_out=.  appreview.proto 

go run main.go --dir_sqls_1=d1 --dir_sqls_2=d2 --dir_sqls_3=d3 -b=d4 -s=sourcevalue
go run main.go --cfg config.yaml
go run main.go sql --cfg=config.yaml

内外穿透方案一：ngrok
https://zhuanlan.zhihu.com/p/1896134392558110239
https://cloud.tencent.com/developer/article/1540655
https://blog.csdn.net/qq_45657541/article/details/147426627

内外穿透方案二：frp自建
https://cloud.tencent.com.cn/developer/article/2417533
https://github.com/fatedier/frp
https://blog.csdn.net/jichencsdn/article/details/138253143
https://zhuanlan.zhihu.com/p/697533940
https://blog.csdn.net/m0_62160083/article/details/144805346


Linux
frps常驻
# 启动frp
sudo systemctl start frps
# 停止frp
sudo systemctl stop frps
# 重启frp
sudo systemctl restart frps
# 查看frp状态
sudo systemctl status frps

# 开机启动
sudo systemctl enable frps

# 修改了.service文件后需要重新加载配置
systemctl daemon-reload

Macos
frpc常驻：

//--------------------- 方法一 ---------------------
# 卸载旧配置（如果存在）
launchctl unload ~/Library/LaunchAgents/com.myapp.frpc.sh.plist 2>/dev/null

# 加载新配置
launchctl load ~/Library/LaunchAgents/com.myapp.frpc.sh.plist

# 立即启动（可选）
launchctl start com.myapp.frpc.sh

//--------------------- 方法二 ---------------------
# 卸载旧配置（如果存在）
launchctl unload ~/Library/LaunchAgents/com.myapp.frpc.plist 2>/dev/null

# 加载新配置
launchctl load ~/Library/LaunchAgents/com.myapp.frpc.plist

# 立即启动（可选）
launchctl start com.myapp.frpc

//------------------------------------------
# 查看状态
launchctl list | grep frp