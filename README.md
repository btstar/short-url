# 短链接服务

## 依赖

    nginx
    redis

## 代码位置

    $GOPATH/zhen.chen/short-url

## 使用

* **生成短链接**

    >curl -X PUT https://lyly.ws/short --data 原始url    
    >生成成功返回短链接 