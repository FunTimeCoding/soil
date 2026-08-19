package gosecret

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strings"
)

func EncodeSecret(path string) (*SecretResult, error) {
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
	decoded, e := ReadDecoded(decodedPath)

	if e != nil {
		return nil, fmt.Errorf("read decoded: %w", e)
	}

	if decoded == nil {
		return nil, nil
	}

	current := make(map[string]string, len(m.Payload))

	for key, value := range m.Payload {
		d, f := base64.StdEncoding.DecodeString(value)

		if f != nil {
			return nil, fmt.Errorf("decode key %s: %w", key, f)
		}

		current[key] = strings.TrimRight(string(d), "\n")
	}

	if reflect.DeepEqual(current, decoded) {
		return &SecretResult{DecodedPath: decodedPath, InSync: true}, nil
	}

	replaced, e := ReplacePayload(b, decoded)

	if e != nil {
		return nil, fmt.Errorf("replace payload: %w", e)
	}

	if e := os.WriteFile(path, replaced, 0644); e != nil {
		return nil, fmt.Errorf("write file: %w", e)
	}

	return &SecretResult{DecodedPath: decodedPath, InSync: false}, nil
}
