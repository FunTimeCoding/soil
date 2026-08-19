package gosecret

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ProcessSecret(
	path string,
	checkMode bool,
) (*SecretResult, error) {
	b, e := os.ReadFile(path)

	if e != nil {
		return nil, fmt.Errorf("read file: %w", e)
	}

	var m SecretManifest

	if e := yaml.Unmarshal(b, &m); e != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", e)
	}

	if m.Kind != "Secret" || len(m.Payload) == 0 {
		return nil, nil
	}

	decodedPath := GetDecodedPath(path)
	decoded := make(map[string]string)

	for key, value := range m.Payload {
		d, e := base64.StdEncoding.DecodeString(value)

		if e != nil {
			return nil, fmt.Errorf("decode key %s: %w", key, e)
		}

		decoded[key] = string(d)
	}

	if checkMode {
		inSync, e := CheckSync(decodedPath, decoded)

		if e != nil {
			return nil, fmt.Errorf("check sync: %w", e)
		}

		return &SecretResult{DecodedPath: decodedPath, InSync: inSync}, nil
	}

	if e := WriteDecoded(decodedPath, decoded); e != nil {
		return nil, fmt.Errorf("write decoded: %w", e)
	}

	return &SecretResult{DecodedPath: decodedPath, InSync: true}, nil
}
