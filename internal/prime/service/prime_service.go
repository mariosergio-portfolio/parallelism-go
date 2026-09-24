package service

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/mycompany/portfolio-go/internal/prime/model"
)

type PrimeService struct{}

func NewPrimeService() *PrimeService {
	return &PrimeService{}
}

func (s *PrimeService) FindPrimes(n, nMaxValue, parallelProcess int) []model.PrimeResult {
	if parallelProcess == 1 {
		return s.findSequential(n, nMaxValue)
	}
	return s.findParallel(n, nMaxValue, parallelProcess)
}

func (s *PrimeService) findSequential(n, nMaxValue int) []model.PrimeResult {
	pid := processId()
	results := make([]model.PrimeResult, 0, n)
	for candidate := 2; candidate <= nMaxValue && len(results) < n; candidate++ {
		if isPrime(candidate) {
			results = append(results, model.PrimeResult{
				Index:     len(results) + 1,
				Value:     candidate,
				ProcessId: pid,
			})
		}
	}
	return results
}

func (s *PrimeService) findParallel(n, nMaxValue, parallelProcess int) []model.PrimeResult {
	primeMap := make(map[int]string, n)
	var mu sync.Mutex
	var stopped atomic.Bool
	var foundCount atomic.Int64
	var wg sync.WaitGroup

	rangeSize := (nMaxValue - 2) / parallelProcess

	for w := 0; w < parallelProcess; w++ {
		from := 2 + w*rangeSize
		to := from + rangeSize - 1
		if w == parallelProcess-1 {
			to = nMaxValue
		}

		wg.Add(1)
		go func(from, to int) {
			defer wg.Done()
			pid := processId()
			for candidate := from; candidate <= to; candidate++ {
				if stopped.Load() {
					return
				}
				if isPrime(candidate) {
					mu.Lock()
					primeMap[candidate] = pid
					mu.Unlock()
					if foundCount.Add(1) >= int64(n) {
						stopped.Store(true)
						return
					}
				}
			}
		}(from, to)
	}

	wg.Wait()

	results := make([]model.PrimeResult, 0, n)
	for candidate := 2; candidate <= nMaxValue && len(results) < n; candidate++ {
		if pid, ok := primeMap[candidate]; ok {
			results = append(results, model.PrimeResult{
				Index:     len(results) + 1,
				Value:     candidate,
				ProcessId: pid,
			})
		}
	}

	return results
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func processId() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	var id int
	fmt.Sscanf(string(buf[:n]), "goroutine %d ", &id)
	return fmt.Sprintf("goroutine-%d", id)
}
