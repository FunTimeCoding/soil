package stamp

type Stamp struct {
	Version    string
	GitHash    string
	CommitDate string
	Module     string
	Dirty      bool
}
