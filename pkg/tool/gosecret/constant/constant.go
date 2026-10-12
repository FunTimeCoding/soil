package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gosecret",
	"Kubernetes secrets: manifest encoding and the gosecretd store",
	"gosecret <command>",
)

type Mode int

const (
	Decode Mode = iota
	Check
	Encode
)

const (
	HostEnvironment     = "GOSECRET_HOST"
	PortEnvironment     = "GOSECRET_PORT"
	InsecureEnvironment = "GOSECRET_INSECURE"
	TokenEnvironment    = "GOSECRET_TOKEN"
)

const (
	DecodeUsage  = "decode"
	CheckUsage   = "check"
	EncodeUsage  = "encode"
	StatusUsage  = "status"
	ListUsage    = "list [path]"
	ShowUsage    = "show <path>"
	RevealUsage  = "reveal <path> <key>"
	PutUsage     = "put <path> <key>"
	MigrateUsage = "migrate <manifest>"
)

const (
	DefaultStore  = "openbao"
	OpaqueType    = "Opaque"
	PathSeparator = "/"
)

const TestManifest = `---
apiVersion: v1
kind: Secret
metadata: {name: example-secret}
type: Opaque
# noinspection SpellCheckingInspection
data:
  ALFA: b25l
  CHARLIE: dGhyZWU=
`
