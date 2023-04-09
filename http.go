package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/index.html")
}

func HTTPServe() {
	go func() {
		http.HandleFunc("/", handler)
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
}
