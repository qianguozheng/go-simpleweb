package airdisk

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/timtadh/data-structures/types"
)

type Upgrade struct {
	Mac string `json:mac`
	Url string `json:url`
	Md5 string `json:md5`
}

type Control struct {
	Mac string `json:mac`
	Switch string `json:switch`
}

type Error struct {
	msg string `json:msg`
}

func UpgradeHandler() http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
		fmt.Println(r.RequestURI)

		r.ParseForm()
		mac := r.Form.Get("mac")
		fmt.Println("mac:", mac)
		if Lupg != nil {
			if Lupg.Has(types.String(mac)){
				upg, err := Lupg.Get(types.String(mac))
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				//if upg.(ConfigUpg) != nil {
					upgrade := Upgrade{
						Mac: upg.(ConfigUpg).cfg.Mac,
						Url: upg.(ConfigUpg).cfg.Url,
						Md5: upg.(ConfigUpg).cfg.Md5,
					}

					data, err := json.Marshal(upgrade)
					if err != nil {
						fmt.Println("Marshal json failed", err.Error())

						goto Failed
					}
					w.Write(data)
					return
				//}
			}
		}

		return
		Failed:
		msgstr := Error{msg:"Upgrade Handler Failed"}
		msg, err := json.Marshal(msgstr)
		if err != nil{
			fmt.Println("err:", err.Error())
		}
		w.Write(msg)
	}
	return http.HandlerFunc(fn)
}

func ControlHandler() http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
		//Do control command
		r.ParseForm()
		mac := r.Form.Get("mac")
		fmt.Println("mac:", mac)
		//Check database to find the mac is set to open Control channel ?

		if Lcfg != nil {
			if (Lcfg.Has(types.String(mac))){
				cfg, err := Lcfg.Get(types.String(mac))

				if err != nil {
					fmt.Println("Error: ", err.Error())
					return
				}


				ctl := Control{Switch:cfg.(ConfigCtl).cfg.Switch,
						Mac:cfg.(ConfigCtl).cfg.Mac}
				data, err := json.Marshal(ctl)
				if err != nil {
					fmt.Println("Control json format failed")
					goto Failed
				}
				w.Write(data)
				return

			}


			return
			Failed:
			msgstr := Error{msg:"Upgrade Handler Failed"}
			msg, err := json.Marshal(msgstr)
			if err != nil {
				fmt.Println("err:", err.Error())
			}
			w.Write(msg)
		}
	}
	return http.HandlerFunc(fn)
}

