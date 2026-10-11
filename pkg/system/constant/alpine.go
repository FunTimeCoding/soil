package constant

const (
	AlpineSignaturePrefix = ".SIGN.RSA."
	AlpinePublicKeySuffix = ".pub"

	AlpineControlDirectory = "control"
	AlpineControlFile      = "control.tar.gz"
	AlpineArchiveFile      = "data.tar.gz"

	AlpineMetadataFile = ".PKGINFO"

	AlpineIndexFile         = "APKINDEX"
	AlpineIndexArchive      = "APKINDEX.tar.gz"
	AlpineIndexName         = "P:"
	AlpineIndexVersion      = "V:"
	AlpineIndexArchitecture = "A:"

	AlpineGzipHeaderSize  = 10
	AlpineGzipMinimumSize = 20

	AlpineArchitecture = "x86_64"
	AlpineRelease      = "1"
)
