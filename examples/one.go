package main

import (
	wbf "../../WebFrame"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.Handle("/api", wbf.HandleHTTP(func(c *wbf.Context) {
		c.
	}))
}
