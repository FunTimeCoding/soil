package constant

const (
	TokenEnvironment        = "ALPINE_TOKEN"
	SignatureKeyEnvironment = "ALPINE_SIGNATURE_KEY"

	SignaturePrefix = ".SIGN.RSA."
	PublicKeySuffix = ".pub"

	RoutePrefix = "apk"

	PackageRoot  = "/apk"
	KeyDirectory = "/key"

	ControlFile = "control.tar.gz"
	ArchiveFile = "data.tar.gz"

	MetadataFile = ".PKGINFO"

	GzipHeaderSize  = 10
	GzipMinimumSize = 20
)
