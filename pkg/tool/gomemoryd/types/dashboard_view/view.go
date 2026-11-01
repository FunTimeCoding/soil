package dashboard_view

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

type View struct {
	Active      int
	Types       map[string]int
	Tags        []*record.TagCount
	Relations   int
	Untyped     int
	Versions    []record.Version
	Impressions []record.Impression
}
