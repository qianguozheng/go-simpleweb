package airdisk

import (
	"net/http"
	"fmt"
)

func Run()  {
	handler := http.NewServeMux()
	handler.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "This is test request")
	})

	handler.Handle("/upgrade", UpgradeHandler())
	handler.Handle("/control", ControlHandler())

	http.ListenAndServe(":9123", handler)
}