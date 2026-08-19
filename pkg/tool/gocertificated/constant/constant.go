package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/identity"
	"os"
	"time"
)

var Identity = identity.New(
	"gocertificated",
	"Internal certificate authority",
	"gocertificated",
).WithInstructions(
	"Internal certificate authority - root and constrained intermediates, certificate signing requests, revocation.",
)

const (
	RootValidityYear         = 20
	IntermediateValidityYear = 10
	LeafValidityDay          = 90
	RevocationListDay        = 7
	ExpiryHorizonDay         = 30
	ClockSkew                = time.Hour
)

const (
	SerialBit  = 128
	SerialBase = 16
)

const (
	RootAuthority  = "root"
	PublishBranch  = "main"
	PublishMessage = "publish"
)

const AnyName = ""
const (
	CertificateBlock    = "CERTIFICATE"
	KeyBlock            = "PRIVATE KEY"
	RevocationListBlock = "X509 CRL"
	SigningRequestBlock = "CERTIFICATE REQUEST"
)

const (
	CertificateFile = "certificate.pem"
	KeyFile         = "key.pem"
)

const (
	CertificateMode os.FileMode = 0644
	KeyMode         os.FileMode = 0600
)

const (
	SecretCertificateKey = "tls.crt"
	SecretKeyKey         = "tls.key"
	SecretType           = "kubernetes.io/tls"
	SecretExtension      = ".yaml"
)

const (
	AuthorityDirectoryEnvironment = "AUTHORITY_DIRECTORY"
	SecretAuthorityEnvironment    = "SECRET_AUTHORITY"
	SecretPathEnvironment         = "SECRET_PATH"
)

const (
	ListAuthorities    = "list_authorities"
	CreateAuthority    = "create_authority"
	GetAuthority       = "get_authority"
	ListCertificates   = "list_certificates"
	IssueCertificate   = "issue_certificate"
	GetCertificate     = "get_certificate"
	RevokeCertificate  = "revoke_certificate"
	SignRequest        = "sign_request"
	PendingPublication = "pending_publication"
	Publish            = "publish"
	RootCertificate    = "root_certificate"
	RevocationList     = "revocation_list"
)

const (
	AuthorityParameter     = "authority"
	KindParameter          = "kind"
	CommonNameParameter    = "common_name"
	CountryParameter       = "country"
	ProvinceParameter      = "province"
	OrganizationParameter  = "organization"
	DomainParameter        = "permitted_domain"
	AddressParameter       = "permitted_address"
	HostParameter          = "host"
	SerialParameter        = "serial"
	RequestParameter       = "request"
	ValidYearParameter     = "valid_year"
	ValidDayParameter      = "valid_day"
	RevokedParameter       = "revoked"
	ExpiresBeforeParameter = "expires_before"
)

const (
	DashboardTitle        = "Dashboard"
	DashboardPath         = "/"
	AuthoritiesTitle      = "Authorities"
	AuthoritiesPath       = "/authorities"
	CertificatesTitle     = "Certificates"
	CertificatesPath      = "/certificates"
	CreateAuthorityTitle  = "Create authority"
	CreateAuthorityPath   = "/authorities/create"
	IssueCertificateTitle = "Issue certificate"
	IssueCertificatePath  = "/certificates/issue"
	PublishTitle          = "Publish"
	PublishPath           = "/publish"
	RootTitle             = "Root certificate"
	RootPath              = "/root"
)

const (
	CommonNameLabel   = "common name"
	CountryLabel      = "country"
	ProvinceLabel     = "province"
	OrganizationLabel = "organization"
)

var (
	ErrorNotFound = errors.New("not found")
	ErrorConflict = errors.New("already live")
)

const (
	QueryFail           = "Failed to query"
	CreateAuthorityFail = "Failed to create the authority"
	IssueFail           = "Failed to issue the certificate"
	SignFail            = "Failed to sign the request"
	RevokeFail          = "Failed to revoke the certificate"
	PublishFail         = "Failed to publish"
	RevocationListFail  = "Failed to build the revocation list"
	AuthorityMissing    = "No live authority of that name"
	CertificateMissing  = "No certificate with that serial"
	AuthorityLive       = "An authority of that name is already live"
	RootMissing         = "No root exists yet"
)

const (
	FixtureProject           = 1
	FixtureClusterAuthority  = "cluster"
	FixtureCountry           = "XX"
	FixtureProvince          = "Example Province"
	FixtureOrganization      = "Example"
	FixtureRootCommonName    = "Example Root CA"
	FixtureIssuingCommonName = "Example Issuing CA"
	FixtureHost              = "service.example.org"
	FixtureCommonName        = "service"
	FixtureDomain            = "example.org"
	FixtureInternalDomain    = "internal"
	FixtureLocalDomain       = "local"
	FixtureAddress           = "192.0.2.0/24"
	FixturePermittedHost     = "192.0.2.10"
	FixtureForeignDomain     = "service.example.net"
	FixtureForeignHost       = "198.51.100.1"
	FixtureLocalHost         = "service.local"
	FixtureRequestHost       = "host.example.org"
	FixtureImpostor          = "impostor"
	FixtureAuthorityDirectory = "certificate"
	FixtureSecretPath        = "manifest/authority-secret.yaml"
)
