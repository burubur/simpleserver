package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8090"
	log.Printf("starting simple http server on port %s", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte(`{"code": "1000", "message":"welcome","data":{}, "error":null}`))
		if err != nil {
			log.Panic("ahahay something error", err)
		}
	})
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Println("ahahay something error", err)
	}

}
