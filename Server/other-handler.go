package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h *OtherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var head string //fisrt subrouter
	head, r.URL.Path = ShiftPath(r.URL.Path)

	switch head {
	case "":
		h.handleRoot().ServeHTTP(w, r)
		return
	case "hello":
		h.handleHello().ServeHTTP(w, r)
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func (h *OtherHandler) handleRoot() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dt := time.Now()
		var message = "Welcome | Other: " + dt.Format("01-02-2006 15:04:05")

		jsonBytes, _ := json.Marshal(&message)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})
}

func (h *OtherHandler) handleHello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dt := time.Now()
		time.Sleep(time.Second * 3)
		var message = "Hello: " + dt.Format("01-02-2006 15:04:05")
		jsonBytes, _ := json.Marshal(&message)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})
}
