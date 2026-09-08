package proxmox

import (
	"context"
	"github.com/luthermonson/go-proxmox"
	"time"
)

type Client struct {
	context    context.Context
	client     *proxmox.Client
	user       string
	password   string
	token      string
	secret     string
	untrusted bool
	log       bool
	verbose   bool
	port      int
	timeout   time.Duration
}
