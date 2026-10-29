package signature

import (
	"crypto/rsa"
	"os"
)

func SignPackageWithKey(
	apkPath string,
	privateKey *rsa.PrivateKey,
	keyName string,
) error {
	unsigned, e := os.ReadFile(apkPath)

	if e != nil {
		return e
	}

	controlEnd, f := FindGzipBoundary(unsigned)

	if f != nil {
		return f
	}

	control := unsigned[:controlEnd]
	packageContent := unsigned[controlEnd:]
	signature, g := SignControl(control, privateKey)

	if g != nil {
		return g
	}

	signed := append(CreateSignatureSegment(signature, keyName), control...)
	signed = append(signed, packageContent...)

	return os.WriteFile(apkPath, signed, 0o644)
}
