package stall_client

import "time"

func (b *Builder) Ping(
	send time.Duration,
	wait time.Duration,
) *Builder {
	b.sendPing = send
	b.ping = wait

	return b
}
