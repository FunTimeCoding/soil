package publish

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (p *Publisher) deliversSecret(r *record.Record) bool {
	return p.secretPath != "" && p.secretAuthority == r.Name
}
