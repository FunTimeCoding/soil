package board

type Connection struct {
	Prometheus Target `yaml:"prometheus"`
	Nextcloud  Target `yaml:"nextcloud"`
	Argocd     Target `yaml:"argocd"`
}
