package airdisk

import (
	"fmt"
	"net/http"
	"encoding/json"
	"../sqlite"
)

type PostBody struct {
	Mac string `json:mac`
}


func UpgradeHandler() http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
		fmt.Println(r.RequestURI)

		var body []byte
		body = make([]byte, 4096)

		n, err := r.Body.Read(body)
		if err!= nil && err.Error() != "EOF"{
			fmt.Println("err", err.Error())
			return
		}
		body[n] = 0

		var mac PostBody

		err = json.Unmarshal(body[:n], &mac)
		if err != nil{
			fmt.Println(err.Error())
		}

		sqlite.DoJob(mac.Mac, 0, w)
	}
	return http.HandlerFunc(fn)
}

func ControlHandler() http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){

		fmt.Println(r.RequestURI)
		var body []byte
		body = make([]byte, 4096)

		n, err := r.Body.Read(body)
		if err!= nil && err.Error() != "EOF"{
			fmt.Println("err", err.Error())
			return
		}
		body[n] = 0

		var mac PostBody

		err = json.Unmarshal(body[:n], &mac)
		if err != nil{
			fmt.Println(err.Error())
		}

		sqlite.DoJob(mac.Mac, 1, w)
	}
	return http.HandlerFunc(fn)
}

