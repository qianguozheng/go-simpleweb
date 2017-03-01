# Simple Web Server Implemented with Golang

## Introduction

This is an example web server that works with wifidog client in embeded device, which only support the http request as heartbeat or control usage.
you can modify it under you convience.

Adding support for wechat wifi connect support.
To basic process need to be go through, now the device not registered.

## API

1. /upgrade - support remote upgrade protocol
2. /control - support remote open ngrok client, which enable ssh login device
3. /wx (at 80 port) - wechat server will use get to verify the server exist.
4. /wx (at 80 port) - wechat server will post event message when something happen.
5. /index.html or / return the portal page.
6. use timer to auto get the access token to server, and do some config.
7. support json config file, when config file specified, it will not use default config from flag.
8. support sqlite3 to store ssid-shopid pair, and userinfo's openid-mac-wechatno-wanmac.

## TODO

1. support add device to support wechat connect wifi.
2. use bootstrap adminlte to add such device metioned in #1



