package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/otaviopontes/api-go/src/config"
	"github.com/otaviopontes/api-go/src/router"
	"github.com/rs/cors"
)

// func init() {
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringBase64)
// }

// @title POSTLOGS API Docs
// @description Server to manage user and posts.
// @version 1.0.0
// @contact.name Otávio Pontes
// @contact.url https://www.otaviopontes.com
// @contact.email otavio.pontes1103@gmail.com
// @BasePath /api
func main() {
	config.Load()

	r := router.Generate()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{config.FrontEndUrl},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "PATCH", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowCredentials: true,
	})

	// Insert the middleware
	handler := cors.Handler(r)

	fmt.Printf("Listening to port %d", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handler))
}
