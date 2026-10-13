package stall_client

import "time"

func (b *Builder) Timeout(d time.Duration) *Builder {
	b.timeout = d

	return b
}
