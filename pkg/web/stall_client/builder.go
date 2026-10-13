package stall_client

import "time"

type Builder struct {
	insecure bool
	long     bool
	timeout  time.Duration
	verbose  bool
	basic    bool
	user     string
	password string
}
