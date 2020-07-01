package mux

import "net/http"

type customeHandler func(w http.ResponseWriter, r *http.Request)

type MuxPropio struct {
	rutasPropias map[string]customeHandler
}

func (this *MuxPropio) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fn, ok := this.rutasPropias[r.URL.Path]; ok {
		fn(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (this *MuxPropio) AddFunc(ruta string, fn customeHandler) {
	this.rutasPropias[ruta] = fn
}

func (this *MuxPropio) AddHandle(ruta string, handle http.Handler) {
	this.rutasPropias[ruta] = handle.ServeHTTP
}

func CreateMux() *MuxPropio {
	newMap := make(map[string]customeHandler)
	return &MuxPropio{newMap}
}
