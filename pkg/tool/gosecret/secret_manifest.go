package gosecret

type SecretManifest struct {
	Kind     string            `yaml:"kind"`
	Metadata *ManifestMetadata `yaml:"metadata"`
	Type     string            `yaml:"type"`
	Payload  map[string]string `yaml:"data"`
}
