package main

import (
	authRoutes "github.com/araquach/apiAuth/routes"
	financeRoutes "github.com/araquach/apiFinance23/routes"
	"github.com/araquach/apiHelpers/middleware"
	teamRoutes "github.com/araquach/apiTeam/routes"
	timeRoutes "github.com/araquach/apiTime/routes"
	db "github.com/araquach/dbService"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5174"}, // replace with your domain
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // or other methods you need
		AllowedHeaders:   []string{"*"},
	})

	// Load API Routes
	authRouter := authRoutes.AuthRouter()
	teamRouter := teamRoutes.TeamRouter()
	financeRouter := financeRoutes.FinanceRouter()
	timeRouter := timeRoutes.TimeRouter()

	mainRouter := mux.NewRouter()
	mainRouter.Use(middleware.ContentTypeMiddleware)
	mainRouter.Use(c.Handler)

	mainRouter.PathPrefix("/api/auth").Handler(authRouter)
	mainRouter.PathPrefix("/api/team").Handler(teamRouter)
	mainRouter.PathPrefix("/api/finance").Handler(financeRouter)
	mainRouter.PathPrefix("/api/time").Handler(timeRouter)

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":"+port, forceSsl(mainRouter))
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
