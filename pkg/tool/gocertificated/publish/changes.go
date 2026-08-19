package publish

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"path"
)

func (p *Publisher) Changes(v []record.Record) ([]*Change, error) {
	var result []*Change

	for _, r := range v {
		reason := join.Space(r.Kind, r.CommonName)
		result = append(
			result,
			NewChange(
				path.Join(
					constant.PublishDirectory,
					r.Name,
					constant.CertificateFile,
				),
				reason,
				r.Certificate,
			),
			NewChange(
				path.Join(constant.PublishDirectory, r.Name, constant.KeyFile),
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
