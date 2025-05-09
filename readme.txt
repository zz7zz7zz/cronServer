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