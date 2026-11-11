package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

const PollInterval = 5 * time.Minute

var Identity = identity.New(
	"gobackrestd",
	"pgBackRest backup exporter",
	"gobackrestd [flags]",
)

const BackrestPeer = "pgbackrest command"
const (
	StanzaLabel = "stanza"
	TypeLabel   = "type"
)

const (
	StatusMetric         = "pgbackrest_stanza_status"
	StatusHelp           = "Stanza status code from pgbackrest info, 0 is ok"
	CompletionMetric     = "pgbackrest_backup_completion_timestamp_seconds"
	CompletionHelp       = "Unix time the latest backup of a type finished"
	SizeMetric           = "pgbackrest_backup_size_bytes"
	SizeHelp             = "Database size covered by the latest backup of a type"
	RepositorySizeMetric = "pgbackrest_backup_repository_size_bytes"
	RepositorySizeHelp   = "Repository size of the latest backup of a type, with the backups it references"
	ErrorMetric          = "pgbackrest_backup_error"
	ErrorHelp            = "1 when the latest backup of a type found page checksum errors"
)
