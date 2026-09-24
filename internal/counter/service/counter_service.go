package service

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mycompany/portfolio-go/internal/counter/model"
)

var goroutineCounter atomic.Int64

type CounterService struct{}

func NewCounterService() *CounterService {
	return &CounterService{}
}

func (s *CounterService) Count(n, countDelay, parallelProcess int) []model.Counter {
	if parallelProcess == 1 {
		return s.countSequential(n, countDelay)
	}
	if parallelProcess == -1 {
		return s.countParallelLikeJavaVirtualThread(n, countDelay)
	}
	return s.countParallel(n, countDelay, parallelProcess)
}

func (s *CounterService) countSequential(n, countDelay int) []model.Counter {
	pid := processId()
	results := make([]model.Counter, n)
	for i := 1; i <= n; i++ {
		time.Sleep(time.Duration(countDelay) * time.Millisecond)
		results[i-1] = model.NewCounter(i, time.Now(), pid)
	}
	return results
}

func (s *CounterService) countParallel(n, countDelay, parallelProcess int) []model.Counter {
	results := make([]model.Counter, n)
	var wg sync.WaitGroup

	chunkSize := (n + parallelProcess - 1) / parallelProcess

	for i := 0; i < parallelProcess; i++ {
		start := i*chunkSize + 1
		end := start + chunkSize - 1
		if end > n {
			end = n
		}
		wg.Add(1)
		go func(from, to int) {
			defer wg.Done()
			pid := processId()
			for number := from; number <= to; number++ {
				time.Sleep(time.Duration(countDelay) * time.Millisecond)
				results[number-1] = model.NewCounter(number, time.Now(), pid)
			}
		}(start, end)
	}

	wg.Wait()
	return results
}

func (s *CounterService) countParallelLikeJavaVirtualThread(n, countDelay int) []model.Counter {
	results := make([]model.Counter, n)
	jobs := make(chan int, n)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pid := processId()
			number := <-jobs
			time.Sleep(time.Duration(countDelay) * time.Millisecond)
			results[number-1] = model.NewCounter(number, time.Now(), pid)
		}()
	}

	wg.Wait()
	return results
}

func processId() string {
	return fmt.Sprintf("goroutine-%d", goroutineCounter.Add(1))
}
