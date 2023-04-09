package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	sessionId := r.URL.Query()["session_id"]
	if len(sessionId) != 0 {
		path := fmt.Sprintf("/tmp/%s", sessionId[0])
		w.Header().Set("Content-Type", "application/json")
		if _, err := os.Stat(path); err == nil {
			ck := http.Cookie{
				Name:  "SESSION_ID",
				Value: sessionId[0],
			}
			http.SetCookie(w, &ck)
		}
		http.ServeFile(w, r, path)
		return
	}

	t, _ := template.ParseFiles("./www/index.html")
	s := Session{Id: uuid.New().String()}
	t.Execute(w, s)
}

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionId, _ := r.Cookie("SESSION_ID")
	path := fmt.Sprintf("/tmp/%s", sessionId.Value)
	f, _ := os.ReadFile(path)
	u := User{}
	json.Unmarshal(f, &u)
	t, _ := template.ParseFiles("./www/session.html")
	t.Execute(w, u)
}

func HTTPServe() {
	go func() {
		http.HandleFunc("/", handler)
		http.HandleFunc("/session", sessionHandler)
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
}
