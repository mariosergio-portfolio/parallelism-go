package main

import (
	"log"

	"github.com/mycompany/portfolio-go/internal/config"
	counterhandler "github.com/mycompany/portfolio-go/internal/counter/handler"
	counterservice "github.com/mycompany/portfolio-go/internal/counter/service"
	primehandler "github.com/mycompany/portfolio-go/internal/prime/handler"
	primeservice "github.com/mycompany/portfolio-go/internal/prime/service"
	"github.com/mycompany/portfolio-go/internal/router"
)

func main() {
	cfg := config.Load()

	cService := counterservice.NewCounterService()
	cHandler := counterhandler.NewCounterHandler(cService)

	pService := primeservice.NewPrimeService()
	pHandler := primehandler.NewPrimeHandler(pService)

	r := router.Setup(cHandler, pHandler, cfg.AllowOrigins)

	log.Printf("Starting portfolio-go on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
