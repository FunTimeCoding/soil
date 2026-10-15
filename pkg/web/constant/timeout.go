package constant

import "time"

const (
	DialTimeout               = 10 * time.Second
	HandshakeTimeout          = 10 * time.Second
	ResponseHeaderTimeout     = 30 * time.Second
	LongResponseHeaderTimeout = 30 * time.Minute
	IdleTimeout               = 90 * time.Second
	SendPingTimeout           = 30 * time.Second
	PingTimeout               = 15 * time.Second
	KeepAlive                 = 30 * time.Second
	MaximumIdle               = 100
)
