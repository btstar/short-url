# 短链接服务
## 集成
    需要一些前置条件
    【1、地址必须是无端口号的；2、域名最好映射一下ssh check out方式（参考：http://holys.im/2016/09/20/go-get-in-gitlab/）】
    go get git.smartswarm.org/zhen.chen/short-url

## 依赖

    nginx
    redis

## 代码位置

    $GOPATH/zhen.chen/short-url

## 使用

* **生成短链接**

    >curl -X PUT https://lyly.ws/short --data 原始url    
    >生成成功返回短链接 