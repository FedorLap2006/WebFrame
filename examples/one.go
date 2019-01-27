package main

import (
	"net/http"

	wbf "../../WebFrame"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.Handle("/api", wbf.HandleHTTP(func(c *wbf.Context) {
		c.WriteStringIO("hello world!")
	}))

	http.ListenAndServe(":8080", mux)
}
