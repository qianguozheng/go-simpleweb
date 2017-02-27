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

func makeSign(t int64, ssid, bssid, mac string, shopid int) string {
	md5Ctx := md5.New()
	str :=[]string{AppId,"Extend",fmt.Sprintf("%d", t),strconv.Itoa(shopid), AuthUrl, mac,
		ssid, bssid, SecretKey}

	ss := strings.Join(str, "")
	fmt.Println(ss)
	io.WriteString(md5Ctx, strings.Join(str,""))
	cipherStr := md5Ctx.Sum(nil)
	return fmt.Sprintf("%x", cipherStr)
}

func (portalCtx *PortalCtx) Portal(c echo.Context) error{

	wanmac := c.QueryParam("wanmac")
	bssid := c.QueryParam("bssid")
	usermac := c.QueryParam("usermac")
	ssid := c.QueryParam("ssid")
	fmt.Println("wanmac=", wanmac)
	fmt.Println("ssid=", ssid)
	fmt.Println("bssid=", bssid)
	fmt.Println("usermac=", usermac)
	//url解码也许

	//ShopId, SSID 从公众号获取 关联起来。
	shopId := models.GetShopId(ssid)
	fmt.Println("shopId=", shopId)

	t := time.Now().UnixNano() / 1000000
	wechatParam := WechatParam{
		AppId: AppId,
		Extend: "Extend",
		Timestamp: fmt.Sprintf("%d", int64(t)), //毫秒
		Sign: makeSign(int64(t), ssid, bssid, wanmac, shopId),
		ShopId: strconv.Itoa(shopId), //strconv.Itoa(ShopId),
		AuthUrl: AuthUrl,
		Mac: wanmac, //"00:0C:43:E1:76:2A",  //不确定是哪个mac地址？
		Ssid: ssid, //"-Subway",
		Bssid: bssid, //"84:5D:D7:E1:76:28",
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

// Processing subscribe interface
func (portalCtx *PortalCtx) Subscribe(c echo.Context) error{
	//return c.String(http.StatusOK, "Hello Upgrade")
	body := new(models.Subscribe)
	if err := c.Bind(body); err != nil{
		return err
	}
	fmt.Println("mac=", body.Mac)
	fmt.Println("usermac=", body.UserMac)
	//respJson, err := models.DoJob(body.Mac, UPGRUDE)
	// 检查数据库中是否存在对应的用户mac对应的openId, 单个公众帐号如何与openId绑定？
	// 思路： 根据路由器mac地址， 找到公众帐号，然后判断对应的公众帐号是否存在对应的openId.

	//if err != nil{
	//	fmt.Println(err)
	//}
	//var resp models.SubscribeResponse
	resp := models.SubscribeResponse{Result:"Fail"}
	var i interface{}
	i = resp

	return c.JSON(http.StatusOK, i)
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