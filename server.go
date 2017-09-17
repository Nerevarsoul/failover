package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type req_struct struct {
    Time string
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    w.Header().Set("Content-Type", "application/json")
    var r req_struct
    err := decoder.Decode(&r)
    if err != nil {
        panic(err)
    }
    defer req.Body.Close()
    log.Println(r.Time)
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
