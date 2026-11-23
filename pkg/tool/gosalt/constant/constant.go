package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New("gosalt", "Salt state tree CLI", "gosalt <command>")

const (
	DirectoryArgument   = "directory"
	DirectoryUsage      = "Roots of one salt environment in the master's order, comma separated or repeated; the first root holding a file wins, top.sls included"
	EnvironmentArgument = "environment"
	EnvironmentUsage    = "Environments the master serves, comma separated; top.sls blocks for any other are dead"
	PillarArgument      = "pillar"
	PillarUsage         = "Pillar roots in the master's order; salt:// files their live pillar names count as used"
)

const (
	OrphansDescription = "Report states and files no live top.sls entry reaches. Jinja is read conservatively: every branch counts, and a templated include makes its matches possible, never orphaned."
	WhyDescription     = "Report whether a state is reached, possible or unreached, and the shortest include chain from each top.sls entry that reaches it, dead environments included."
	TraceDescription   = "Report every state and file one top.sls target brings in, across environments, and which of them nothing else would reach once the target is removed."
)
