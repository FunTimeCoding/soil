package constant

var (
	BackupLabels = []string{
		HypervisorLabel,
		TypeLabel,
		IdentifierLabel,
		NameLabel,
	}
	GuestLabels = []string{
		HypervisorLabel,
		NodeLabel,
		TypeLabel,
		IdentifierLabel,
		NameLabel,
	}
	HypervisorLabels = []string{HypervisorLabel}
	NodeLabels       = []string{HypervisorLabel, NodeLabel}
	StorageLabels    = []string{
		HypervisorLabel,
		NodeLabel,
		StorageLabel,
		PluginLabel,
		ContentLabel,
	}
)
