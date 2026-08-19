package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func Certificates(v []record.Record) []server.CertificateResponse {
	result := make([]server.CertificateResponse, 0, len(v))

	for _, r := range v {
		result = append(result, *Certificate(&r))
	}

	return result
}
