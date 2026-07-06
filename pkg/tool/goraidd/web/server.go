package web

import (
	"github.com/funtimecoding/soil/pkg/raid_parser"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	store      *store.Store
	elitePath  string
	outputPath string
	parser     *raid_parser.Client
	view       *view.View
	registry   *palette.Registry
}
