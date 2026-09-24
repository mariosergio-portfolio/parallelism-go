package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mycompany/portfolio-go/internal/counter/model"
	"github.com/mycompany/portfolio-go/internal/counter/service"
)

const maxEstimatedMs = 60_000

type CounterHandler struct {
	service *service.CounterService
}

func NewCounterHandler(svc *service.CounterService) *CounterHandler {
	return &CounterHandler{service: svc}
}

func (h *CounterHandler) Count(requestContext *gin.Context) {
	var req struct {
		N               int `form:"n"               binding:"required,min=1"`
		CountDelay      int `form:"countDelay"      binding:"required,min=1"`
		ParallelProcess int `form:"parallelProcess" binding:"required,min=-1"`
	}

	log.Println("PARALLELISMO GO - Request:", requestContext.Request.URL)

	if err := requestContext.ShouldBindQuery(&req); err != nil {
		requestContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//	estimatedMs := int64(math.Ceil(float64(req.N)/float64(req.ParallelProcess))) * int64(req.CountDelay)
	//	if estimatedMs > maxEstimatedMs {
	//		c.Status(http.StatusNotFound)
	//		return
	//}

	start := time.Now()
	counters := h.service.Count(req.N, req.CountDelay, req.ParallelProcess)
	end := time.Now()

	displayed := counters
	if len(counters) > 101 {
		displayed = append(counters[:100:100], counters[len(counters)-1])
	}

	log.Println("PARALLELISMO GO - Request COMPLETTED:", requestContext.Request.URL)

	requestContext.JSON(http.StatusOK, model.CounterResponse{
		Summary:  model.NewCounterSummary(start, end, req.N, req.CountDelay, req.ParallelProcess),
		Counters: displayed,
	})
}
