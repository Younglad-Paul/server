package main

import (
	"backend/cookie"
	"backend/graph"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "8080"

func withResponseWriter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), cookie.ResponseWriterKey, w)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	// Load env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, proceeding without it: %v", err)
	}
	

	// Get keys from the env file
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect to MongoDB
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %s", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s", err)
	}
	defer client.Disconnect(ctx)

	// Create a new resolver with MongoDB client
	resolver := graph.NewResolver(client)

	// Create GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Set up CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://aramco.vercel.app", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler

	// Set up router
	router := chi.NewRouter()
	router.Use(corsHandler)
	router.Use(withResponseWriter)

	// HTTP handlers
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", withResponseWriter(srv))

	// WebSocket handler

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
