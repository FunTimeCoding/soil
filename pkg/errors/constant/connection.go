package constant

import "regexp"

const (
	Host       = "host"
	Path       = "path"
	Connection = "connection"
)

var QueryPattern = regexp.MustCompile(`(https?://[^\s"?]+)\?[^\s"]*`)
