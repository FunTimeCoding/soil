package publish

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/change"
	"path"
)

func (p *Publisher) Changes(v []record.Record) ([]*change.Change, error) {
	var result []*change.Change

	for _, r := range v {
		reason := join.Space(r.Kind, r.CommonName)
		result = append(
			result,
			change.New(
				path.Join(
					p.authorityDirectory,
					r.Name,
					constant.CertificateFile,
				),
				reason,
				r.Certificate,
			),
			change.New(
				path.Join(p.authorityDirectory, r.Name, constant.KeyFile),
				reason,
				r.Key,
			),
		)

		if !p.deliversSecret(&r) {
			continue
		}

		secret, e := p.secretChange(&r)

		if e != nil {
			return nil, e
		}

		result = append(result, secret...)
	}

	return result, nil
}
