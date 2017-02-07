package airdisk

import (
	"net/http"
	"fmt"
	//"../sqlite"
)


func DBInit(){
	//sqlite.Connect()
}
func Run()  {
	handler := http.NewServeMux()
	handler.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "This is test request")
	})

	handler.Handle("/upgrade", UpgradeHandler())
	handler.Handle("/control", ControlHandler())

	////Config
	//handler.Handle("/config/upgrade", ConfigUpgHandler())
	//handler.Handle("/config/control", ConfigCtlHandler())

	http.ListenAndServe(":80", handler)

}