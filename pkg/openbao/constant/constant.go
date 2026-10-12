package constant

import "time"

const (
	HostEnvironment     = "OPENBAO_HOST"
	PortEnvironment     = "OPENBAO_PORT"
	InsecureEnvironment = "OPENBAO_INSECURE"
	RoleEnvironment     = "OPENBAO_ROLE"
)

const (
	Mount            = "secret"
	MetadataFormat   = "%s/metadata/%s"
	KeysField        = "keys"
	ListParameter    = "list"
	ListMethod       = "LIST"
	TokenPath        = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	RequestTimeout   = 10 * time.Second
	SecretKind       = "secret"
	RenewDenominator = 4
)
