package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"sync"
	"time"
)

type Client struct {
	results  map[string]*query_result.Result
	floats   map[string]float64
	failures map[string]error
	moments  []time.Time
	mutex    sync.Mutex
	Failure  error
	Closed   int
}
