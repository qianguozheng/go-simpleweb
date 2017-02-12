package airdisk

import (
	"net/http"
	"fmt"
	//"../sqlite"
)


func DBInit(){
	//sqlite.Connect()
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	staticHandler := http.FileServer(http.Dir("./template/"))
	staticHandler.ServeHTTP(w, r)
	return
}

func Run()  {
	handler := http.NewServeMux()
	handler.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "This is test request")
	})

	handler.Handle("/upgrade", UpgradeHandler())
	handler.Handle("/control", ControlHandler())

	handler.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//handler.HandleFunc("/", index)
	handler.HandleFunc("/index.html", index)
	////Config
	//handler.Handle("/config/upgrade", ConfigUpgHandler())
	//handler.Handle("/config/control", ConfigCtlHandler())

	http.ListenAndServe(":38001", handler)

}