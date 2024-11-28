package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"td/internal/controller"
	"td/internal/logic"
	"td/internal/middleware"
	"td/internal/respond"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	frontendDistPath := os.Getenv("FRONTEND_PATH")
	if frontendDistPath == "" {
		log.Fatal("$FRONTEND_PATH must be set")
	}

	// Create main router and set up terrible cors
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.LoggingMiddleware)

	g := logic.NewGame()
	gl := logic.NewGameLogic(g)
	ws := logic.NewGameWebSocket(&gl)
	r := respond.NewResponder()
	c := controller.NewGameController(gl, r)

	gameRouter := chi.NewRouter()

	// V1 Routes
	gameRouter.Get("/state", c.State)
	gameRouter.Get("ws", ws.HandleWebSocket)

	// Mount the v1Router under /v1
	router.Mount("/api/game", gameRouter)


	// Serve Frontend
	devMode := os.Getenv("DEV") == "true"

	// Dev
	if devMode {
		viteServerURL := os.Getenv("VITE_SERVER_URL")
		if viteServerURL == "" {
			log.Fatal("$VITE_SERVER_URL must be set")
		}
		proxyURL, err := url.Parse(viteServerURL)
		if err != nil {
			log.Fatalf("Invalid VITE_SERVER_URL: %v", err)
		}
		proxy := httputil.NewSingleHostReverseProxy(proxyURL)

		router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			r.URL.Host = proxyURL.Host
			r.URL.Scheme = proxyURL.Scheme
			r.Host = proxyURL.Host
			proxy.ServeHTTP(w, r)
		})

	// Build
	} else {
		staticHandler := http.FileServer(http.Dir(frontendDistPath))
		router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			filePath := frontendDistPath + r.URL.Path
			_, err := os.Stat(filePath)
			if os.IsNotExist(err) {
				// If file does not exist, serve index.html
				http.ServeFile(w, r, frontendDistPath+"/index.html")
				return
			}
			staticHandler.ServeHTTP(w, r)
		})
	}
	

	// Start Server
	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}