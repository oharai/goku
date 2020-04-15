package main

import (
	"net/http"

	"github.com/garupanojisan/goku/router"
)

func main() {
	r := router.NewRouter()
	r.GET("/", handler).
		Before(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("before"))
		}).
		After(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("after"))
		})
	r.Run(":8080")
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	msg := query.Get("msg")
	num := query.Get("num")
	w.Write([]byte(msg + " " + num))
}
