package main

import (
	"flag"
	"fmt"
	"github.com/azimjohn/jprq/jprq_tcp"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	var baseHost string
	flag.StringVar(&baseHost, "host", "tcp.camelot-register.uz", "Base Host")
	flag.Parse()

	j := jprq_tcp.New(baseHost)
	r := mux.NewRouter()
	r.HandleFunc("/_ws/", j.WebsocketHandler)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	fmt.Println("Server is running on Port 4500")
	log.Fatal(http.ListenAndServe(":4500", r))
}
