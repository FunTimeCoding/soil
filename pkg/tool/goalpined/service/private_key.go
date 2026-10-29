package service

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"os"
)

func (s *Service) privateKey() (*rsa.PrivateKey, error) {
	path := s.keyPath()
	content, e := os.ReadFile(path)

	if e != nil {
		return nil, e
	}

	block, _ := pem.Decode(content)

	if block == nil {
		return nil, unexpected.Format(
			"signature key without PEM block: %s",
			path,
		)
	}

	if key, f := x509.ParsePKCS1PrivateKey(block.Bytes); f == nil {
		return key, nil
	}

	parsed, g := x509.ParsePKCS8PrivateKey(block.Bytes)

	if g != nil {
		return nil, fmt.Errorf("parse signature key: %w", g)
	}

	result, okay := parsed.(*rsa.PrivateKey)

	if !okay {
		return nil, unexpected.Format("signature key is not RSA: %s", path)
	}

	return result, nil
}
