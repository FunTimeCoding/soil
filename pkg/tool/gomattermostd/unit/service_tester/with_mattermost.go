package service_tester

import "github.com/funtimecoding/soil/pkg/chat/face"

func WithMattermost(c face.MattermostClient) Option {
	return func(o *Tester) {
		o.client = c
	}
}
