package model

import (
	"fmt"
	"strconv"
	"time"
)

const TimeFormat = "15:04:05.000Z"

type Counter struct {
	Number          int    `json:"number"`
	CompletedTime   string `json:"completedTime"`
	CompletedTimeMs int64  `json:"completedTimeMs"`
	ProcessId       string `json:"processId"`
}

func NewCounter(number int, completedAt time.Time, processId string) Counter {
	return Counter{
		Number:          number,
		CompletedTime:   completedAt.UTC().Format(TimeFormat),
		CompletedTimeMs: completedAt.UnixMilli(),
		ProcessId:       processId,
	}
}

type SummaryRequest struct {
	N               int    `json:"n"`
	CountDelay      int    `json:"countDelay"`
	ParallelProcess string `json:"parallelProcess"`
}

type SummaryResponse struct {
	StartTime         string `json:"startTime"`
	EndTime           string `json:"endTime"`
	DurationMs        int64  `json:"durationMs"`
	DurationFormatted string `json:"durationFormatted"`
}

type CounterSummary struct {
	Request  SummaryRequest  `json:"request"`
	Response SummaryResponse `json:"response"`
}

func NewCounterSummary(start, end time.Time, n, countDelay, parallelProcess int) CounterSummary {
	durationMs := end.Sub(start).Milliseconds()
	h := durationMs / 3_600_000
	m := (durationMs % 3_600_000) / 60_000
	s := (durationMs % 60_000) / 1_000
	ms := durationMs % 1_000

	var processStr string = strconv.Itoa(parallelProcess)

	if parallelProcess == -1 {
		processStr = fmt.Sprintf("(%s) goroutines=n=%s)", strconv.Itoa(parallelProcess), strconv.Itoa(n))
	}

	return CounterSummary{
		Request: SummaryRequest{
			N:               n,
			CountDelay:      countDelay,
			ParallelProcess: processStr,
		},
		Response: SummaryResponse{
			StartTime:         start.UTC().Format(TimeFormat),
			EndTime:           end.UTC().Format(TimeFormat),
			DurationMs:        durationMs,
			DurationFormatted: fmt.Sprintf("%02d:%02d:%02d:%03d", h, m, s, ms),
		},
	}
}

type CounterResponse struct {
	Summary  CounterSummary `json:"summary"`
	Counters []Counter      `json:"counters"`
}
