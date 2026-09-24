# portfolio-go

Go 1.26 + Gin REST API — Counter feature.

## Requirements

- Go 1.26+ ([https://go.dev/dl/](https://go.dev/dl/))

## Setup

```bash
# Download dependencies
go mod tidy

# Run
go run ./cmd/main.go

# Build
go build -o portfolio-go ./cmd/main.go
```

## Endpoint

```
GET http://localhost:8081/api/counter?n=10&countDelay=200&parallelProcess=4
```

### Validation rules
- `n`, `countDelay`, `parallelProcess` must all be integers ≥ 1
- Estimated execution time `⌈n / parallelProcess⌉ × countDelay` must not exceed **60 000 ms** → returns `404`

### Response

```json
{
  "summary": {
    "startTime":         "10:00:00 000Z",
    "endTime":           "10:00:00 512Z",
    "durationMs":        512,
    "durationFormatted": "00:00:00:512"
  },
  "counters": [
    { "number": 1, "completedTime": "10:00:00 201Z", "threadName": "goroutine-6" },
    { "number": 2, "completedTime": "10:00:00 201Z", "threadName": "goroutine-7" }
  ]
}
```

## Project structure

```
portfolio-go/
├── cmd/
│   └── main.go                          ← entry point
└── internal/
    ├── counter/
    │   ├── model/counter.go             ← domain model + response types
    │   ├── service/counter_service.go   ← business logic (sequential / parallel)
    │   └── handler/counter_handler.go  ← Gin handler
    └── router/router.go                 ← route registration
```
