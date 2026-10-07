package message_meta

import "github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"

type Meta struct {
	Type      string
	Statistic response.Statistic
}
