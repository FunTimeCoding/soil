package constant

import "errors"

const (
	HostEnvironment  = "ATLASSIAN_HOST"
	UserEnvironment  = "ATLASSIAN_USER"
	TokenEnvironment = "ATLASSIAN_TOKEN" // #nosec G101 not a hardcoded secret
)

var ErrorNotFound = errors.New("not found")
