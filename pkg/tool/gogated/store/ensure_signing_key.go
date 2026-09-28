package store

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/signing_key"
	"github.com/google/uuid"
)

func (s *Store) EnsureSigningKey() (*rsa.PrivateKey, string, error) {
	var count int64
	s.mapper.Model(signing_key.Stub()).Count(&count)

	if count > 0 {
		var row signing_key.SigningKey
		e := s.mapper.First(&row).Error

		if e != nil {
			return nil, "", e
		}

		block, _ := pem.Decode([]byte(row.PrivateKey))
		key, e := x509.ParsePKCS1PrivateKey(block.Bytes)
		errors.PanicOnError(e)

		return key, row.Identifier, nil
	}

	key, e := rsa.GenerateKey(rand.Reader, 2048)

	if e != nil {
		return nil, "", e
	}

	identifier := uuid.New().String()
	privateEncoded := string(
		pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(key),
			},
		),
	)
	publicBytes, f := x509.MarshalPKIXPublicKey(&key.PublicKey)
	errors.PanicOnError(f)
	publicEncoded := string(
		pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicBytes}),
	)
	e = s.mapper.Create(
		signing_key.New(
			identifier,
			privateEncoded,
			publicEncoded,
			constant.AlgorithmRS256,
		),
	).Error

	if e != nil {
		return nil, "", e
	}

	return key, identifier, nil
}
