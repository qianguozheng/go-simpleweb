package airdisk

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Upgrade struct {
	Mac string `json:mac`
	Url string `json:url`
	Md5 string `json:md5`
}

type Control struct {
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
		if mac == "12345678"{
			upg := Upgrade{
				Mac: "12345678",
				Url: "http://127.0.0.1/test.bin",
				Md5: "1234567890123456789012",
			}

			data, err := json.Marshal(upg)
			if err != nil{
				fmt.Println("Marshal json failed", err.Error())

				goto Failed
			}
			w.Write(data)
			return
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
		if mac == "12345678" {
			//type Control struct {
			//	Switch string `json:switch`
			//}
			ctl := Control{Switch:"On"}
			data, err := json.Marshal(ctl)
			if err != nil{
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
		if err != nil{
			fmt.Println("err:", err.Error())
		}
		w.Write(msg)
	}
	return http.HandlerFunc(fn)
}

