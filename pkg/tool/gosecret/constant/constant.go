package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gosecret",
	"Kubernetes secret encoder and checker",
	"gosecret [--check|--encode] [--directory <path>]",
)

type Mode int

const (
	Decode Mode = iota
	Check
	Encode
)

const TestManifest = `---
apiVersion: v1
kind: Secret
metadata: {name: example-secret}
type: Opaque
# noinspection SpellCheckingInspection
data:
  ALPHA: b25l
  GAMMA: dGhyZWU=
`
