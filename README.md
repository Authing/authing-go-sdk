# SDK for Go

Authing Go SDK 目前支持 Golang 1.8+ 版本。

GitHub 地址：[https://github.com/authing/authing-go-sdk](https://github.com/authing/authing-go-sdk)。

## 安装

```bash
go get github.com/authing/authing-go-sdk
```

## 快速开始

```go
package main

import (
	"authing-golang-sdk/lib/management"
	"authing-golang-sdk/lib/model"
	"log"
)

const (
	userPoolId  = "60a6f97f3f50c7a9483e313d"
	appSecret = "d254623f808ba850d5d5ea7b07bead60"
)

func main() {
	client := management.NewClient(userPoolId, appSecret)
	log.Println("==========通过 ID 获取用户信息==========")
	resp, err := client.Detail("60a6f9ad5bcccc51834950c5")
	log.Printf("%+v\n", resp)
}
```

[如何获取 UserPool ID 和 UserPool Secret ？](https://docs.authing.cn/v2/guides/faqs/get-userpool-id-and-secret.html)

## API 使用实例

### 通过 ID 获取用户信息

通过用户 ID 获取用户详情

```go
client := management.NewClient(userPoolId, appSecret)
resp, _ := client.Detail("60a6f9ad5bcccc51834950c5")
```

### 

