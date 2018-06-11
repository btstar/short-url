# 短链接服务
## 集成到现有项目

    import "github.com/czsilence/short-url/server"

    server.Init()

## 单独运行

    go get github.com/czsilence/short-url
    cd $GOPATH/src/github.com/czsilence/short-url
    go run main.go

## 依赖

    redis

## 使用

* **生成短链接**

    >curl -X PUT http://path_to_server_root/short --data 原始url    
    >生成成功返回短链接 