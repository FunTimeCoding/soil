package package_server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/linux/alpine/signature"
)

func signPackage(
	apkPath string,
	keyName string,
) {
	k, e := loadPrivateKey(keyName)
	errors.PanicOnError(e)
	signature.SignPackageWithKey(apkPath, k, keyName)
}
