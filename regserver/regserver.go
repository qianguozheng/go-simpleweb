package regserver

import (
	"fmt"
	"strings"
	"net/http"
	"sort"
	"crypto/sha1"
	"io"
	"github.com/labstack/echo"
	"../models"
)

var Token string //"12345678901234567890qwertyuioqgz"
type WeChatCtx struct {

}

func NewWeChatCtx()*WeChatCtx  {
	wc := WeChatCtx{}
	return &wc
}

func makeSignature(timestamp, nonce string) string{
	//将token, timestamp, nonce三个参数进行字典序排序
	s1 := []string{Token, timestamp, nonce}
	sort.Strings(s1)
	//将三个参数字符串拼接成一个字符串进行sha1加密
	s := sha1.New()
	io.WriteString(s, strings.Join(s1,""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func (wx WeChatCtx) Handle(c echo.Context) error {
	fmt.Println("--- WeChat Serve ---")

	//提取signature, timestamp, nonce, echostr
	sig := c.QueryParam("signature")
	time := c.QueryParam("timestamp")
	nonce := c.QueryParam("nonce")
	echostr := c.QueryParam("echostr")

	fmt.Println("sig=", sig)
	fmt.Println("time=", time)
	fmt.Println("nonce=", nonce)
	fmt.Println("echostr=", echostr)
	//token := "shit"

	//验证hashcode == signature

	signatureGen := makeSignature(time, nonce)

	if signatureGen != sig {
		return c.String(http.StatusOK,"")
	} else {
		return c.String(http.StatusOK, echostr)
	}
}

//<xml>
//<ToUserName><![CDATA[toUser]]></ToUserName>
//<FromUserName><![CDATA[FromUser]]></FromUserName>
//<CreateTime>123456789</CreateTime>
//<MsgType><![CDATA[event]]></MsgType>
//<Event><![CDATA[subscribe]]></Event>
//</xml>
type UnsubscribeEvent struct {
	ToUserName string	`xml:"ToUserName"`
	FromUserName string	`xml:"FromUserName"`
	CreateTime int64	`xml:"CreateTime"`
	MsgType string		`xml:"MsgType"`
	Event string		`xml:"Event"`
}
func (wx WeChatCtx) HandlePost(c echo.Context) error {
	fmt.Println("--- WeChat Serve ---")

	//提取signature, timestamp, nonce, echostr
	sig := c.QueryParam("signature")
	time := c.QueryParam("timestamp")
	nonce := c.QueryParam("nonce")
	openid := c.QueryParam("openid")

	fmt.Println("sig=", sig)
	fmt.Println("time=", time)
	fmt.Println("nonce=", nonce)
	fmt.Println("openid=", openid)
	//token := "shit"

	//验证hashcode == signature
	body := new(UnsubscribeEvent)

	if err := c.Bind(body); err != nil{
		return err
	}
	signatureGen := makeSignature(time, nonce)

	//if signatureGen != sig {
	//	return c.String(http.StatusOK,"")
	//} else {
	//	return c.String(http.StatusOK, openid)
	//}
	fmt.Println(signatureGen)

	fmt.Println("ToUserName=",body.ToUserName)
	fmt.Println("FromUserName=", body.FromUserName)
	fmt.Println("CreateTime=", body.CreateTime)
	fmt.Println("MsgType=", body.MsgType)
	fmt.Println("Event=", body.Event)

	switch  {
	case strings.HasPrefix(body.Event, "unsubscribe"):
	//更新数据库，删除未订阅事件
		models.RemoveUserInfo(body.FromUserName, body.ToUserName)
	case strings.HasPrefix(body.Event, "subscribe"):
		models.AddWechatNo2UserInfo(body.FromUserName, body.ToUserName)
	default:
		fmt.Println("Not known")
	}

	return c.String(http.StatusOK, "")
}