package openbao

import (
	"github.com/openbao/openbao/api/v2"
	"sync"
	"time"
)

type Client struct {
	client    *api.Client
	role      string
	tokenPath string
	expiry    time.Time
	lock      sync.Mutex
}
