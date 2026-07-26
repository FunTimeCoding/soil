package mattermost

import "errors"

var (
	ErrorNotConfigured = errors.New("not configured")
	ErrorNotFound      = errors.New("not found")
)
