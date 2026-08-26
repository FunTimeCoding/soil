package package_server

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func loadPrivateKey(keyName string) (*rsa.PrivateKey, error) {
	keyPath := filepath.Join(constant.KeyDirectory, keyName)
	keyBytes := system.ReadBytesUnsafe(keyPath)
	block, _ := pem.Decode(keyBytes)

	if block == nil {
		return nil, validation.New("no PEM block in %s", keyPath)
	}

	if key, f := x509.ParsePKCS1PrivateKey(block.Bytes); f == nil {
		return key, nil
	}

	keyInterface, f := x509.ParsePKCS8PrivateKey(block.Bytes)

	if f != nil {
		return nil, fmt.Errorf("parse private key: %w", f)
	}

	rsaKey, okay := keyInterface.(*rsa.PrivateKey)

	if !okay {
		return nil, validation.New("private key is not RSA")
	}

	return rsaKey, nil
}
