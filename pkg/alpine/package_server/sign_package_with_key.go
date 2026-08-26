package package_server

import (
	"crypto/rsa"
	"github.com/funtimecoding/soil/pkg/system"
)

func SignPackageWithKey(
	apkPath string,
	privateKey *rsa.PrivateKey,
	keyName string,
) {
	unsigned := system.ReadBytesUnsafe(apkPath)
	controlEnd := FindGzipBoundary(unsigned)
	control := unsigned[:controlEnd]
	packageContent := unsigned[controlEnd:]
	signature := SignControl(control, privateKey)
	signatureSegment := CreateSignatureSegment(signature, keyName)
	signed := append(signatureSegment, control...)
	signed = append(signed, packageContent...)
	system.WriteFile(apkPath, signed, 0644)
}
