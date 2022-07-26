## Drone telegram plugin

The result of the plug-in is a message that the bot will send you.:
```
Repo name

✅ 第1次构建成功
🕙 耗时：1秒
📖 提交分支：master
🎅 提交者：hongzx
🔗 详情：https://ci.your.site/service/1
📃 提交信息： your commit
```

Variables
  - *proxy_url* - You can use any proxy tool if api telegram is not available from your country(do not fill out to keep default) Example format : *https://your.tg.proxy*
  - *token* - Your telegram bot token - Required
  - *chat_id* - Chat ID, which will be sent to the bot notifications - Required

Example pipeline
```yml
kind: pipeline
name: project-go-api

steps:
  - name: build
    image: golang:latest
    pull: if-not-exists
    environment:
      GOPROXY: "https://goproxy.cn,direct" 
    volumes:
      - name: pkgdeps
        path: /go/pkg
    commands:
      - CGO_ENABLED=0 go build -o project-go-api
      
  - name: telegram
    image: hongzhuangxian/telegram-drone-plugin
    settings:
      proxy_url: "https://your.proxy.url"
      token:
        from_secret: telegram_token
      chat_id:
        from_secret: telegram_chat_id
```
Build packed:

    set GOOS=linux
    set GOARCH=amd64
    set CGO_ENABLED=0
    go build -o telegram-drone-plugin

Build image:

    docker build -t hongzhuangxian/telegram-drone-plugin .

Push image:

    docker push hongzhuangxian/telegram-drone-plugin
