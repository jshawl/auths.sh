package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	sessionId := r.URL.Query()["session_id"]
	if len(sessionId) != 0 {
		w.Header().Set("Content-Type", "application/json")
		// if file exists
		// write cookie
		// cookie1 := &http.Cookie{Name: "sample", Value: "sample", HttpOnly: false}
		// http.SetCookie(w, cookie1)
		http.ServeFile(w, r, fmt.Sprintf("/tmp/%s", sessionId[0]))
		return
	}

	t, _ := template.ParseFiles("./www/index.html")
	s := Session{Id: uuid.New().String()}
	t.Execute(w, s)
}

func HTTPServe() {
	go func() {
		http.HandleFunc("/", handler)
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
}
