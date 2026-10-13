package proxmox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/luthermonson/go-proxmox"
	"time"
)

type Client struct {
	context   context.Context
	client    *proxmox.Client
	user      string
	password  string
	token     string
	secret    string
	untrusted bool
	log       bool
	verbose   bool
	port      int
	timeout   time.Duration
	pool      *pool.Pool
}
