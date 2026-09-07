package search_option

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"

type Option struct {
	Query      string
	Limit      int
	Collection string
	Full       bool
	Mode       string
	Metadata   map[string]string
	Reranker   face.Reranker
	Exclude    []string
}
