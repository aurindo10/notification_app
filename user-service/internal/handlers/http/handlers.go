package server

import "net/http"

type Handlers struct {
	mux *http.ServeMux
}

func (c *Handlers) StartHandlers() {
	c.mux.HandleFunc("GET /registeruser", func(w http.ResponseWriter, r *http.Request) {
		res := []byte("Olá")
		w.Write(res)
	})
}

func NewHandlers(mux *http.ServeMux) *Handlers {
	return &Handlers{
		mux: mux,
	}
}
