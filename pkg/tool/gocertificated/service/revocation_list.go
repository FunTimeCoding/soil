package service

import (
	"crypto/x509"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"math/big"
)

func (s *Service) RevocationList(name string) ([]byte, error) {
	a, e := s.authorityOf(name)

	if e != nil {
		return nil, e
	}

	f := store.NewFilter()
	f.Authority = name
	f.Revoked = new(true)
	revoked, g := s.store.Certificates(f)

	if g != nil {
		return nil, g
	}

	var entry []x509.RevocationListEntry

	for _, r := range revoked {
		serial, okay := new(big.Int).SetString(r.Serial, constant.SerialBase)

		if !okay {
			return nil, errors.Format("serial is not a number", r.Serial)
		}

		entry = append(
			entry,
			x509.RevocationListEntry{
				SerialNumber:   serial,
				RevocationTime: *r.Revoked,
			},
		)
	}

	return a.RevocationList(entry), nil
}
