package airdisk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	hashtable "github.com/timtadh/data-structures/hashtable"
	"github.com/timtadh/data-structures/types"
)

//type Upgrade struct {
//	Mac string `json:mac`
//	Url string `json:url`
//	Md5 string `json:md5`
//}
//
//type Control struct {
//	Switch string `json:switch`
//}

const (
	CONTROL = iota
	UPGRADE
)

type ConfigCtl struct {
	cfg Control
	kind int
}
type ConfigUpg struct {
	cfg Upgrade
	kind int
}

/*
	POST /config/control
	{
		"mac":"xxx",
		"switch":"on"
	}
	POST /config/upgrade
	{
		"mac":"xxx",
		"url":"xxx",
		"md5":"xxx",
		"version":"xxx"
	}
*/

//var Lcfg *utils.HashTable
//var Lupg *utils.HashTable
//var test *hashtable.Hash
var Lcfg *hashtable.Hash
var Lupg *hashtable.Hash

func ConfigUpgHandler() http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
		fmt.Println(r.RequestURI)
		if Lupg == nil {
			//Lupg = utils.New(1000)
			Lupg = hashtable.NewHashTable(100)
			fmt.Println("Init Hash Table")
		}
		//var body []byte
		var cfgupg ConfigUpg
		//r.Body().Read(body)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil{
			fmt.Fprintf(w, "Read body failed")
		}
		err = json.Unmarshal(body, &cfgupg.cfg)
		if err != nil{
			fmt.Fprintf(w, "Set Failed: %s", err.Error())
			return
		}
		cfgupg.kind = UPGRADE

		Lupg.Put(types.String(cfgupg.cfg.Mac), cfgupg)
		fmt.Fprintf(w, "OK")

	}
	return http.HandlerFunc(fn)
}

func ConfigCtlHandler() http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){

		if Lcfg == nil {
			//Lcfg = utils.New(1000)
			Lcfg = hashtable.NewHashTable(100)
			fmt.Println("Init Hash Table")
		}

		var cfgctl ConfigCtl
		//r.Body().Read(body)
		body, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, &cfgctl.cfg)
		if err != nil{
			fmt.Fprintf(w, "Set Failed: %s", err.Error())
		}
		cfgctl.kind = CONTROL
		fmt.Println(cfgctl.cfg.Mac)
		Lcfg.Put(types.String(cfgctl.cfg.Mac), cfgctl)
		fmt.Fprintf(w,"OK")
	}
	return http.HandlerFunc(fn)
}



