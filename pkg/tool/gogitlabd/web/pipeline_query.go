package web

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"net/url"
	"strconv"
)

func pipelineQuery(
	instance string,
	project int64,
	pipeline int64,
) string {
	return join.Ampersand(
		[]string{
			key_value.Equals(
				constant.InstanceParameter,
				url.QueryEscape(instance),
			),
			key_value.Equals(
				argumentConstant.Project,
				strconv.FormatInt(project, 10),
			),
			key_value.Equals(
				argumentConstant.Pipeline,
				strconv.FormatInt(pipeline, 10),
			),
		},
	)
}
