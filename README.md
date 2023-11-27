# 新功能
### 1. 增加了Google搜索框，方便快速搜索
### 2. 增加了背景图，为了使文字清晰点，背景图加了毛玻璃的效果
使用方法：把背景图命名为bg.jpg，然后放在主机上要映射到docker的目录

![homepage screenshot](https://github.com/luckylu/flare/raw/main/screenshots/homepage.jpg)

docker-compose.yml
```shell
version: '3.6'

services:
  flare:
    image: luckylu/flare
    restart: always
    # 默认无需添加任何参数，如有特殊需求
    # 可阅读文档 https://github.com/soulteary/docker-flare/blob/main/docs/advanced-startup.md
    command: flare
    # 启用账号登陆模式
    # command: flare --nologin=0
    # environment:
      # 如需开启用户登陆模式，需要先设置 `nologin` 启动参数为 `0`
      # 如开启 `nologin`，未设置 FLARE_USER，则默认用户为 `flare`
      # - FLARE_USER=flare
      # 指定你自己的账号密码，如未设置 `FLARE_USER`，则会默认生成密码并展示在应用启动日志中
      # - FLARE_PASS=your_password
      # 是否开启“使用向导”，访问 `/guide`
      # - FLARE_GUIDE=1
    ports:
      - 5005:5005
    volumes:
      - /flare:/app
```

# Flare

Challenge all bookmarking apps and websites directories, Aim to Be a best performance monster.

🚧 **Code is being prepared and refactored, commits are slow.**

## Feature

**Simple**, **Fast**, **Lightweight** and super **Easy** to install and use.

- Written in Go (Golang) and a little Modern vanilla Javascript only.
- Doesn't depend any database or any complicated framework.
- Single executable, no dependencies required, good docker support.
- You can choose whether to enable various functions according to your needs: offline mode, weather, editor, account, and so on.

## ScreenShot

TBD

## Documentation

TBD

- Browse automatically generated program documentation:
    - `godoc --http=localhost:8080`

