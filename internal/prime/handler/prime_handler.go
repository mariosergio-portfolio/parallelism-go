package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mycompany/portfolio-go/internal/counter/model"
	primemodel "github.com/mycompany/portfolio-go/internal/prime/model"
	"github.com/mycompany/portfolio-go/internal/prime/service"
)

type PrimeHandler struct {
	service *service.PrimeService
}

func NewPrimeHandler(svc *service.PrimeService) *PrimeHandler {
	return &PrimeHandler{service: svc}
}

func (h *PrimeHandler) FindPrimes(c *gin.Context) {
	var req struct {
		N               int `form:"n"               binding:"required,min=1"`
		NMaxValue       int `form:"nMaxValue"       binding:"required,min=2"`
		ParallelProcess int `form:"parallelProcess" binding:"required,min=1"`
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start := time.Now()
	primes := h.service.FindPrimes(req.N, req.NMaxValue, req.ParallelProcess)
	end := time.Now()

	displayed := primes
	if len(primes) > 101 {
		displayed = append(primes[:100:100], primes[len(primes)-1])
	}

	durationMs := end.Sub(start).Milliseconds()
	startStr := start.UTC().Format(model.TimeFormat)
	endStr := end.UTC().Format(model.TimeFormat)

	summary := primemodel.PrimeSummary{
		Request: primemodel.PrimeSummaryRequest{
			N:               req.N,
			NMaxValue:       req.NMaxValue,
			ParallelProcess: req.ParallelProcess,
		},
		Response: primemodel.NewPrimeSummaryResponse(startStr, endStr, durationMs, len(primes)),
	}

	c.JSON(http.StatusOK, primemodel.PrimeResponse{
		Summary: summary,
		Primes:  displayed,
	})
}
