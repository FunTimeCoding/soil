package store

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (s *Store) Certificates(f *Filter) ([]record.Record, error) {
	var result []record.Record
	d := s.database.Order("not_after")

	if f.Authority != "" {
		d = d.Where("issuer = ?", f.Authority)
	}

	if f.Kind != "" {
		d = d.Where("kind = ?", f.Kind)
	}

	if f.Before != nil {
		d = d.Where("not_after < ?", *f.Before)
	}

	d = revocationClause(d, f.Revoked)

	if f.Limit > 0 {
		d = d.Limit(f.Limit)
	}

	return result, d.Find(&result).Error
}
