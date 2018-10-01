package main

import (
	"net/http"

	"github.com/QO-Development/cashFlowApp/controllers/httpHandlers"
)

func main() {
	//HTML handlers
	http.HandleFunc("/", httpHandlers.Home)

	//File handlers
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	//Start the server
	http.ListenAndServe(":8024", nil)
}
