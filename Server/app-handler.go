package main

import (
	"encoding/json"
	"net/http"
)

func (h *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//rootPath := r.URL.Path

	var head string //fisrt subrouter
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch head {
	case "":
		h.handleHome().ServeHTTP(w, r)
		return
	case "other":
		h.OtherHandler.ServeHTTP(w, r)
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func (h *App) handleHome() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var message = "Welcome | Server"
		jsonBytes, _ := json.Marshal(&message)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})
}
