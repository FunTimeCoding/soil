package dependency

type Dependency struct {
	MonitorIdentifier string
	Path              string
	Version           string
	Directory         string
	Files             []string
	Identifiers       []string
	Family            string
	concern           []string
}
