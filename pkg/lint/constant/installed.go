package constant

const (
	StaleBinaryKey       = "stale_binary"
	StaleBinaryText      = "Installed binary is older than the latest tag - in %s run gobuild --copy-to-bin %s"
	UnresolvedBinaryKey  = "unresolved_binary"
	UnresolvedBinaryText = "Installed binary carries no version - in %s run gobuild --copy-to-bin %s"

	CommandDirectory = "cmd"
)
