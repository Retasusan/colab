package main

import (
	"log"
	"net/http"

	"github.com/Retasusan/colab_backend/internal/db"
	"github.com/Retasusan/colab_backend/internal/org"
)

func main() {
	gdb, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	orgHandler := org.NewHandler(gdb)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/api/orgs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			orgHandler.ListOrganizations(w, r)
		case http.MethodPost:
			orgHandler.CreateOrganization(w, r)
		default:
			http.Error(w, "method not allowd", http.StatusMethodNotAllowed)
		}
	})

	log.Println("server started :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
