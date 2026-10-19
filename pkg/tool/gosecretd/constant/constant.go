package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"gosecretd",
	"Infrastructure secrets kept in OpenBao",
	"gosecretd",
).WithInstructions(
	"Infrastructure secrets in OpenBao, one path per Kubernetes Secret: <namespace>/<secret-name>, its keys the Secret's own keys. Values stay masked; reveal_secret returns one key per call. put_secret merges keys into a secret and never removes any.",
)

const (
	Status       = "status"
	ListSecrets  = "list_secrets"
	ShowSecret   = "show_secret"
	RevealSecret = "reveal_secret"
	PutSecret    = "put_secret"
)

const (
	PathParameter   = "path"
	ValuesParameter = "values"
)

const (
	KeyKind       = "secret key"
	Folder        = "/"
	ParentSegment = ".."
)

const (
	WatchInterval  = time.Minute
	SealedGauge    = "openbao_sealed"
	SealedHelp     = "Whether OpenBao reports itself sealed, 1 when sealed."
	ReachableGauge = "openbao_reachable"
	ReachableHelp  = "Whether OpenBao answered its health check, 1 when it did."
)
