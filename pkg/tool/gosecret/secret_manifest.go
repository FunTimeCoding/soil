package gosecret

type SecretManifest struct {
	Kind    string            `yaml:"kind"`
	Payload map[string]string `yaml:"data"`
}
