package service

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"math/big"
)

func publicKeyToSigningKey(
	key *rsa.PublicKey,
	keyIdentifier string,
) *SigningKey {
	return &SigningKey{
		KeyType:       constant.KeyTypeRSA,
		Use:           constant.KeyUseSigning,
		KeyIdentifier: keyIdentifier,
		Algorithm:     constant.AlgorithmRS256,
		N:             base64.RawURLEncoding.EncodeToString(key.N.Bytes()),
		E: base64.RawURLEncoding.EncodeToString(
			big.NewInt(int64(key.E)).Bytes(),
		),
	}
}
