package startup

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
)

type Service struct {
	roster     *roster.Service
	transcript *transcript.Service
	search     *search.Service
}
