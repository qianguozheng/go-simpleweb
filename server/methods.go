package airdisk

import (
	"github.com/labstack/echo"
	"net/http"
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
	return c.String(http.StatusOK, "Hello Upgrade")
}

func (portalCtx *PortalCtx) Config(c echo.Context) error{
	return c.String(http.StatusOK, "Hello Config")
}
