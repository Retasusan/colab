package main

import (
	"log"
	"net/http"

	appauth "github.com/Retasusan/colab_backend/internal/auth"
	"github.com/Retasusan/colab_backend/internal/db"
	"github.com/Retasusan/colab_backend/internal/org"
)

func main() {
	gdb, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	authMiddleware, err := appauth.NewMiddleware("http://localhost:3000/api/auth/jwks")
	if err != nil {
		log.Fatal(err)
	}

	orgHandler := org.NewHandler(gdb)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.Handle("/api/orgs", authMiddleware.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			orgHandler.ListOrganizations(w, r)
		case http.MethodPost:
			orgHandler.CreateOrganization(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})))
	log.Println("server started :8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(mux)))
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
