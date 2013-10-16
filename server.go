package main

import (
	"net/http"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte("{\"message\":\"hello world\"}"))
}

func main() {
	servemux := http.NewServeMux()
   servemux.HandleFunc("/", baseHandler)

   http.ListenAndServe("localhost:4321", servemux)
}
