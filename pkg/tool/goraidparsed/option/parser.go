package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Parser struct {
	option.Web
	ParserPath   string
	TemplatePath string
	OutputPath   string
	TopStats     func(path string) face.TopStatsParser
}
