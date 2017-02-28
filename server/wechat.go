package airdisk

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"io/ioutil"
	"errors"
	"github.com/bitly/go-simplejson"
	"../models"
)

var (
	AppId string
	AppSecret string
	SecretKey string
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	Expires_In int `json:"expires_in"`
}

type ErrorStr struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}
func GetAccessToken() (*AccessToken, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		AppId,
		AppSecret,
	)
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("Get AccessToken failed", err.Error())
		return nil, errors.New("Http get failed")
	}
	defer resp.Body.Close()
	//var data []byte
	//data = make([]byte, 4096)
	//resp.Body.Read(data)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("read response failed")
		return nil, errors.New("ReadAll utils failed")
	}

	var at AccessToken
	err = json.Unmarshal(data, &at)
	if err != nil{
		fmt.Println("json failed", err.Error())
		var es ErrorStr
		err = json.Unmarshal(data, &es)
		if err != nil{
			fmt.Println("json failed2", err.Error())
			return nil, errors.New("Unmarshal json failed")
		}

	}
	fmt.Println("data:", string(data))
	fmt.Println("access_token=", at.AccessToken)

	return &at, nil
}
var accesstoken string
//添加Portal型设备
//https://api.weixin.qq.com/bizwifi/apportal/register?access_token=ACCESS_TOKEN

type PortalDevice struct {
	ShopId int `json:"shop_id"`
	Ssid string `json:"ssid"`
	Reset bool `json:"reset"`
}
type DataResponse struct {
	SecretKey string `json:"secretkey"`
}
type PortalResponse struct {
	Errcode int `json:"errcode"`
	Data DataResponse `json:"data"`
}
func RegisterPortalDev(shopId int, ssid string) (int, error){
	url := fmt.Sprintf("https://api.weixin.qq.com/bizwifi/apportal/register?access_token=%s", accesstoken)
	//err := http.Post(url, json.Marshal(&PortalDevice{ShopId:123, Ssid:"hello", Reset:false}))
	postData , err := json.Marshal(PortalDevice{ShopId:shopId, Ssid:ssid, Reset:false})
	resp, err:= http.Post(url, "application/json",
		strings.NewReader(string(postData)))

	if err != nil{
		fmt.Println("http Post failed", err.Error())
		return -1, errors.New("Http post failed")
	}
	defer resp.Body.Close()

	data, err:= ioutil.ReadAll(resp.Body)

	if err != nil{
		fmt.Println("http post read failed", err.Error())
		return -2, errors.New("Read from http post failed")
	}

	var pr PortalResponse
	err = json.Unmarshal(data, &pr)
	if err != nil{
		fmt.Println("Parse json failed", err.Error())
		return -3, errors.New("Parse json failed")
	}

	fmt.Println("Total Response: ", string(data))
	return pr.Errcode, nil
}


type PageList struct {
	PageIndex int `json:"pageindex"`
	PageSize int `json:"pagesize"`
}
type PageDataRecords struct {
	ShopId int `json:"shop_id"`
	ShopName string `json:"shop_name"`
	Ssid string `json:"ssid"`
	ProtocolType int `json:"protocol_type"`
	Sid string `json:"sid"`
}
type PageData struct {
	TotalCount int `json:"totalcount"`
	PageIndex int `json:"pageindex"`
	PageCount int `json:"pagecount"`
	Records PageDataRecords `json:"records"`
}
type PageResp struct {
	ErrCode int `json:"errcode"`
	Data PageData `json:"data"`
}


func GetShopList() ([]models.ShopIdSSID, error){
	//https://api.weixin.qq.com/bizwifi/shop/list?access_token=ACCESS_TOKEN
	url := fmt.Sprintf("https://api.weixin.qq.com/bizwifi/shop/list?access_token=%s", accesstoken)
	postData, err := json.Marshal(PageList{PageIndex:1, PageSize:10})
	resp, err := http.Post(url, "application/json",
		strings.NewReader(string(postData)))

	if err != nil{
		fmt.Println("Http post failed", err.Error())
		return nil, errors.New("Http post failed")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("Http post read response failed")
		return nil, errors.New("http post read response failed")
	}

	fmt.Println(string(data))


	var ss []models.ShopIdSSID
	cJson,_ := simplejson.NewJson(data)
	record,_:= cJson.Get("data").Get("records").Array()
	for _, i := range record{
		newRecord, _ := i.(map[string]interface{})
		shopId := newRecord["shop_id"]
		shopName := newRecord["shop_name"]
		ssid := newRecord["ssid"]
		sid := newRecord["sid"]

		fmt.Println("shopid=", shopId)
		fmt.Println("shopName=", shopName)
		fmt.Println("ssid=", ssid)
		fmt.Println("sid=", sid)

		var tmpSs models.ShopIdSSID
		tmpSs.Ssid = ssid.(string)
		x,_ := shopId.(json.Number).Int64()
		tmpSs.ShopId = int(x)

		ss = append(ss, tmpSs)
	}

	//fmt.Println("sss=", ss)
	//
	//for _, v:= range ss{
	//	fmt.Println(v.Ssid)
	//	fmt.Println(v.ShopId)
	//}

	return ss, nil
}


type Shop struct {
	ShopId int `json:"shop_id"`
}
func GetShopWiFiInfo(shopId int)  {
	url := fmt.Sprintf("https://api.weixin.qq.com/bizwifi/shop/get?access_token=%s", accesstoken)
	postData ,err := json.Marshal(Shop{ShopId:shopId})
	resp, err := http.Post(url, "application/json", strings.NewReader(string(postData)))
	if err != nil{
		fmt.Println("http post failed", err.Error())
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("http post read failed", err.Error())
	}
	fmt.Println("Data=", string(data))
	return
}

//Test usage
//func main(){
//
//	at, err := GetAccessToken()
//	if err != nil{
//		fmt.Println("Get AccessToken failed")
//		return
//	}
//	accesstoken = at.AccessToken
//	fmt.Println("Expire:", at.Expires_In, "AccessToken:", at.AccessToken)
//
//	GetShopList()
//	GetShopWiFiInfo(4177281)
//}


/// Currently, It just update the token peridicoly
/// 两个小时更新一次token值
var at *AccessToken

//TODO: secret key need to be update by different SSID/ShopId
func Routines(){
	var err error
	at, err = GetAccessToken()
	if err != nil{
		fmt.Println("Get AccessToken failed")
		return
	}

	fmt.Println("Expire:", at.Expires_In, "AccessToken:", at.AccessToken)
	accesstoken = at.AccessToken
	shopSsid, _ := GetShopList()

	for _, v := range shopSsid{
		//ssid := v.Ssid
		//shopId := v.ShopId

		if (false == models.CheckExistence(v)){
			models.InsertSSID(v)
		}
	}
}