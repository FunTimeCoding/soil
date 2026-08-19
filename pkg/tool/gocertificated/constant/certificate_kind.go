package constant

type CertificateKind string

const (
	KindRoot         CertificateKind = "root"
	KindIntermediate CertificateKind = "intermediate"
	KindServer       CertificateKind = "server"
	KindClient       CertificateKind = "client"
)

var AuthorityKind = map[CertificateKind]bool{
	KindRoot:         true,
	KindIntermediate: true,
}
