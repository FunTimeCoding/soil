package label

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/gazetteer"

func New(c gazetteer.Source) *Collector {
	return &Collector{netbox: c}
}
