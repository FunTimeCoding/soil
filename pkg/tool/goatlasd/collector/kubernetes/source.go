package kubernetes

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"

func (c *Collector) Source() string {
	return constant.SourceKubernetes
}
