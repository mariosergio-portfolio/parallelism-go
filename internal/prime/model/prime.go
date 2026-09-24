package model

import "fmt"

type PrimeResult struct {
	Index     int    `json:"index"`
	Value     int    `json:"value"`
	ProcessId string `json:"processId"`
}

type PrimeSummaryRequest struct {
	N               int `json:"n"`
	NMaxValue       int `json:"nMaxValue"`
	ParallelProcess int `json:"parallelProcess"`
}

type PrimeSummaryResponse struct {
	StartTime         string `json:"startTime"`
	EndTime           string `json:"endTime"`
	DurationMs        int64  `json:"durationMs"`
	DurationFormatted string `json:"durationFormatted"`
	TotalFound        int    `json:"totalFound"`
}

type PrimeSummary struct {
	Request  PrimeSummaryRequest  `json:"request"`
	Response PrimeSummaryResponse `json:"response"`
}

type PrimeResponse struct {
	Summary PrimeSummary  `json:"summary"`
	Primes  []PrimeResult `json:"primes"`
}

func NewPrimeSummaryResponse(startTime, endTime string, durationMs int64, totalFound int) PrimeSummaryResponse {
	h := durationMs / 3_600_000
	m := (durationMs % 3_600_000) / 60_000
	s := (durationMs % 60_000) / 1_000
	ms := durationMs % 1_000
	return PrimeSummaryResponse{
		StartTime:         startTime,
		EndTime:           endTime,
		DurationMs:        durationMs,
		DurationFormatted: fmt.Sprintf("%02d:%02d:%02d:%03d", h, m, s, ms),
		TotalFound:        totalFound,
	}
}
