package router

import (
	counterhandler "github.com/mycompany/portfolio-go/internal/counter/handler"
	primehandler "github.com/mycompany/portfolio-go/internal/prime/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(counterHandler *counterhandler.CounterHandler, primeHandler *primehandler.PrimeHandler, allowOrigins []string) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		api.GET("/counter", counterHandler.Count)
		api.GET("/prime", primeHandler.FindPrimes)
	}

	return r
}
