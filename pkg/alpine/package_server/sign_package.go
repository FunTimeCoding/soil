package package_server

import "github.com/funtimecoding/soil/pkg/errors"

func signPackage(
	apkPath string,
	keyName string,
) {
	k, e := loadPrivateKey(keyName)
	errors.PanicOnError(e)
	SignPackageWithKey(apkPath, k, keyName)
}
