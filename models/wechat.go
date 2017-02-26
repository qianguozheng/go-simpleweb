package models

import (
	"fmt"
)

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

	// shopId, SSID

	// 根据ssid，查找shopId, 填写认证参数

 */
type ShopIdSSID struct {
	ShopId int
	Ssid string
}

func CheckExistence(ss ShopIdSSID) bool {
	if ss.Ssid != "" && ss.ShopId != 0 {
		rows, err := db.Query("select * from wechat where ssid=$1 and shopid=$2", ss.Ssid, ss.ShopId)
		if err != nil{
			fmt.Println(err.Error())
			return false
		}

		var ssid string
		var shopId int

		if (rows.Next()) {
			err = rows.Scan(&ssid, &shopId)
		}

		rows.Close()

		if ssid != "" && shopId != 0{
			return true
		}

	} else {
		return true
	}
	return false
}

func InsertSSID(ss ShopIdSSID){
	_, err := db.Exec("insert into wechat (shopid, ssid) values ($1, $2)", ss.Ssid, ss.ShopId)
	if err != nil{
		fmt.Println(err.Error())
	}
}

func GetShopId(ssid string) int {
	//if ssid != ""{
	//	rows, err := db.Query("select * from wechat where ssid=$1", ssid)
	//
	//	if err != nil{
	//		fmt.Println(err.Error())
	//		return 0
	//	}
	//
	//	var shopId int
	//	if (rows.Next()){
	//		err = rows.Scan(&ssid, &shopId)
	//	}
	//	rows.Close()
	//
	//	return shopId
	//}
	//return 0
	return 4177281
}