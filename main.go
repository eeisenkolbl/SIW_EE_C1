// @title           Residencial Guaraní API
// @version         1.0
// @description     API REST para la landing page de Residencial Guaraní. Incluye chatbot FAQ con árbol de decisión, departamentos y servicios.
// @contact.name    Ernesto Eisenkolbl
// @license.name    MIT
// @host            localhost:8080
// @BasePath        /
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/eeisenkolbl/SIW_EE_C1/docs"
	"github.com/eeisenkolbl/SIW_EE_C1/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	}))

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Human-readable docs
	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/docs.html")
	})

	// API
	r.Post("/api/chat", handlers.PostChat)
	r.Get("/api/departamentos", handlers.GetDepartamentos)
	r.Get("/api/servicios", handlers.GetServicios)
	r.Get("/api/respuestas", handlers.GetRespuestas)
	r.Get("/api/respuestas/{id}", handlers.GetRespuestaByID)

	// Static files
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/*", fs)

	fmt.Println("Servidor iniciado en http://localhost:8080")
	fmt.Println("Swagger UI en   http://localhost:8080/swagger/")
	log.Fatal(http.ListenAndServe(":8080", r))
}
