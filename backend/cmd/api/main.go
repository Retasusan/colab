package main

import (
	"log"
	"net/http"

	"github.com/Retasusan/colab_backend/internal/db"
)

func main() {
	_, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("server started :8080")
	http.ListenAndServe(":8080", nil)
}
