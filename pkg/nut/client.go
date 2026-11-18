package nut

import "time"

type Client struct {
	address string
	timeout time.Duration
}
