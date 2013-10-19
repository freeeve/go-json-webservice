package main

import (
	"encoding/json"
	"github.com/daaku/go.httpgzip"
	"net/http"
)

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func baseHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"hello world\"}"))
}

func userCreateHandler(w http.ResponseWriter, req *http.Request) {
	ucr := UserCreateRequest{}
	json.NewDecoder(req.Body).Decode(&ucr)
	CreateUser(ucr)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"success\"}"))
}

func main() {
	servemux := http.NewServeMux()
	servemux.HandleFunc("/", baseHandler)
	servemux.HandleFunc("/user-create", userCreateHandler)

	http.ListenAndServeTLS("localhost:4321", "cert.pem", "key.pem", httpgzip.NewHandler(servemux))
}
