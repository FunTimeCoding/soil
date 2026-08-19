package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
)

func Authority(r *record.Record) *server.AuthorityResponse {
	c := r.Material().Certificate
	result := &server.AuthorityResponse{
		Name:        r.Name,
		Kind:        server.AuthorityKind(r.Kind),
		Serial:      r.Serial,
		CommonName:  r.CommonName,
		Certificate: r.Certificate,
		NotBefore:   r.Start,
		NotAfter:    r.End,
		Published:   r.Published != nil,
	}

	if d := c.PermittedDNSDomains; len(d) > 0 {
		result.PermittedDomain = &d
	}

	if a := address(c.PermittedIPRanges); len(a) > 0 {
		result.PermittedAddress = &a
	}

	return result
}
