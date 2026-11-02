package option

import "github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"

type Parser struct {
	Address       string
	ParserPath    string
	TemplatePath  string
	OutputPath    string
	ServiceTokens []string
	TopStats      func(path string) face.TopStatsParser
}
