package signature

import (
	"crypto/rsa"
	"github.com/funtimecoding/soil/pkg/errors"
)

func MustSignPackageWithKey(
	apkPath string,
	privateKey *rsa.PrivateKey,
	keyName string,
) {
	errors.PanicOnError(SignPackageWithKey(apkPath, privateKey, keyName))
}
