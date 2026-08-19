package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func Authorities(v []record.Record) []server.AuthorityResponse {
	result := make([]server.AuthorityResponse, 0, len(v))

	for _, r := range v {
		result = append(result, *Authority(&r))
	}

	return result
}
