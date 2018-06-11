# 短链接服务
## 集成
    需要项目权限
    go get git.smartswarm.org:3000/zhen.chen/short-url.git

## 依赖

    nginx
    redis

## 代码位置

    $GOPATH/zhen.chen/short-url

## 使用

* **生成短链接**

    >curl -X PUT https://lyly.ws/short --data 原始url    
    >生成成功返回短链接 