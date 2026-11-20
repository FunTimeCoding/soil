package grant

import "github.com/funtimecoding/soil/pkg/tool/goclauded/face"

type Grant struct {
	client       face.Client
	session      string
	directory    string
	repositories []string
	loaded       bool
	failure      error
	notes        []string
}
