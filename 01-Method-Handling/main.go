// Creating a Http Server
package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

// ServeHTTP implements http.Handler.
func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return
		default:
			log.Fatal()
			w.Write([]byte("404 not found"))
			return
		}
	case "POST":
		switch r.URL.Path {
		case "/":
		}
	}
}

func main() {
	api := &api{addr: ":8080"}
	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}

	srv.ListenAndServe()

}
