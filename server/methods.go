package airdisk

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"../sqlite"
)

type WechatParam struct {
	AppId string
	Extend string
	Timestamp string
	Sign string
	ShopId string
	AuthUrl string
	Mac string
	Ssid string
	Bssid string
}

type PortalCtx struct {

}

// Unified Protocol

type ControlReq struct {
	Mac string `json:mac`
}
type UpgradeReq struct {
	Mac string `json:mac`
	Ver string `json:version`
}

const (
	UPGRUDE = iota
	CONTROL
)

func NewPortalCtx() *PortalCtx{
	lc := PortalCtx{}

	return &lc
}

func (portalCtx *PortalCtx) Portal(c echo.Context) error{
	//c.SetHandler(UpgradeHandler)
	//return c.String(http.StatusOK, "Hello Portal")
	wechatParam := WechatParam{
		AppId: "AppId",
		Extend: "Extend",
		Timestamp: "Timestamp",
		Sign: "Sign",
		ShopId: "ShopId",
		AuthUrl: "AuthUrl",
		Mac: "aa:bb:cc:dd:ee:ff",
		Ssid: "Hello",
		Bssid: "HelloBssid",
	}
	return c.Render(http.StatusOK, "WechatParam", wechatParam)
}

func (portalCtx *PortalCtx) Upgrade(c echo.Context) error{
	//return c.String(http.StatusOK, "Hello Upgrade")
	body := new(UpgradeReq)
	if err := c.Bind(body); err != nil{
		return err
	}
	fmt.Println(body.Mac)
	respJson, err := sqlite.DoJob(body.Mac, UPGRUDE)
	if err != nil{
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, respJson)
}


func (portalCtx *PortalCtx) Control(c echo.Context) error{

	body := new(ControlReq)
	if err := c.Bind(body); err != nil{
		return err
	}
	fmt.Println(body.Mac)
	respJson, err := sqlite.DoJob(body.Mac, CONTROL)
	if err != nil{
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, respJson)
	//return c.String(http.StatusOK, "Hello Config")

}
