// Creating a Http Server
package main

import "net/http"

type server struct {
	addr string
}

// ServeHTTP implements http.Handler.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	s := &server{addr: ":8080"}
	http.ListenAndServe(s.addr, s)
}
