package main

import (
	"github.com/araquach/apiAuth/routes"
	finance "github.com/araquach/apiFinance23/routes"
	apiTime "github.com/araquach/apiTime/routes"
	db "github.com/araquach/dbService"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db.DBInit(dsn)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Load API Routes
	finance.FinanceRouter()
	apiTime.TimeRouter()

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":"+port, forceSsl(&routes.R))
}

func forceSsl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("GO_ENV") == "production" {
			if r.Header.Get("x-forwarded-proto") != "https" {
				sslUrl := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, sslUrl, http.StatusTemporaryRedirect)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
