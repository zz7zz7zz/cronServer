protoc --go_out=.  appreview.proto 

go run main.go --dir_sqls_1=d1 --dir_sqls_2=d2 --dir_sqls_3=d3 -b=d4 -s=sourcevalue
go run main.go --cfg config.yaml
go run main.go sql --cfg=config.yaml