package sentry

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic"

type Client struct {
	basic        *basic.Client
	organization string
	projects     []string
}
