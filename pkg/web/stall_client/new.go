package stall_client

import "github.com/funtimecoding/soil/pkg/web/constant"

func New() *Builder {
	return &Builder{
		sendPing: constant.SendPingTimeout,
		ping:     constant.PingTimeout,
	}
}
