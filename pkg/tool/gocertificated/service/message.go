package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func message(v []record.Record) string {
	var name []string

	for _, r := range v {
		name = append(name, r.Name)
	}

	return join.Space(constant.PublishMessage, join.Comma(name))
}
