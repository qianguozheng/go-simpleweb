package airdisk

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"../models"
	"time"
	"crypto/md5"
	"io"
	"strings"
	"strconv"
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

const (
	//AppId = ""
	//Extend
	ShopId = 4177281
	AuthUrl = "http://hiweeds.net:38001/auth"
	//SecretKey = "685aec96360b737c175b13343cc53388"

)
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

func makeSign(t int64) string {
	md5Ctx := md5.New()
	str :=[]string{AppId,"Extend",fmt.Sprintf("%d", t),strconv.Itoa(ShopId), AuthUrl, "00:0C:43:E1:76:2A",
		"-Subway", "84:5D:D7:E1:76:28", SecretKey}

	ss := strings.Join(str, "")
	fmt.Println(ss)
	io.WriteString(md5Ctx, strings.Join(str,""))
	cipherStr := md5Ctx.Sum(nil)
	return fmt.Sprintf("%x", cipherStr)
}

func (portalCtx *PortalCtx) Portal(c echo.Context) error{

	//wanmac := c.QueryParam("mac")
	//bssid := c.QueryParam("bssid")
	//usermac := c.QueryParam("user_mac")

	//url解码也许

	//填入如下参数


	t := time.Now().UnixNano() / 1000000
	wechatParam := WechatParam{
		AppId: AppId,
		Extend: "Extend",
		Timestamp: fmt.Sprintf("%d", int64(t)), //毫秒
		Sign: makeSign(int64(t)),
		ShopId: strconv.Itoa(ShopId),
		AuthUrl: AuthUrl,
		Mac: "00:0C:43:E1:76:2A",  //不确定是哪个mac地址？
		Ssid: "-Subway",
		Bssid: "84:5D:D7:E1:76:28",
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
	respJson, err := models.DoJob(body.Mac, UPGRUDE)
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
	respJson, err := models.DoJob(body.Mac, CONTROL)
	if err != nil{
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, respJson)
	//return c.String(http.StatusOK, "Hello Config")

}


func (portalCtx *PortalCtx) Auth(c echo.Context) error{
	extend := c.QueryParam("extend")
	openId := c.QueryParam("openId")
	tid := c.QueryParam("tid")

	fmt.Println("Extend=", extend)
	fmt.Println("OpenId=", openId)
	fmt.Println("Tid=", tid)

	return c.String(http.StatusOK, "")
	//return c.String(http.StatusOK, "Hello Config")
}