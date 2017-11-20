package simpleserver

import (
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hellow world"))
} 

func init() {
	http.HandleFunc("/", handle)
}
