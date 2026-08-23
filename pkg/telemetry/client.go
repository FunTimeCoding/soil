package telemetry

import (
	"net/http"
	"sync"
)

type Client struct {
	base   string
	client *http.Client
	group  sync.WaitGroup
}
