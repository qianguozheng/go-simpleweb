package models

// RouterMac | shopid | openid |

/***
整理需求：

	框架：
	注册服务器80端口 http://hiweeds.net/wx
		1. 注册时需要验证
		2. 认证时需要向这个服务器请求放行与否
	路由器认证服务器 38001
		1. 与路由器通信，做一些控制之类的。
		2. 提供回调接口，获取OpenId与mac地址绑定。 hiweeds.net:38001/auth

	SSID固定，BSSID可能会验证。

	路由器Mac地址 --- BSSID, SSID,

	OpenId, WechatNo?,


 */
