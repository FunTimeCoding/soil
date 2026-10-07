package service

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service/response"
	"math/big"
)

func publicKeyToSigningKey(
	key *rsa.PublicKey,
	keyIdentifier string,
) *response.SigningKey {
	return response.NewSigningKey(
		constant.KeyTypeRSA,
		constant.KeyUseSigning,
		keyIdentifier,
		constant.AlgorithmRS256,
		base64.RawURLEncoding.EncodeToString(key.N.Bytes()),
		base64.RawURLEncoding.EncodeToString(big.NewInt(int64(key.E)).Bytes()),
	)
}
