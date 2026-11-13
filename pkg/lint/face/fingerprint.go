package face

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"maps"
	"slices"
)

func (s *Set) Fingerprint() string {
	lines := make(map[string]bool)

	for _, all := range s.methodInterfaces {
		for _, i := range all {
			for method, signature := range i.Methods {
				lines[join.Space(i.Package, i.Name, method, signature)] = true
			}
		}
	}

	h := sha256.New()

	for _, l := range slices.Sorted(maps.Keys(lines)) {
		h.Write([]byte(l))
		h.Write([]byte{0})
	}

	return hex.EncodeToString(h.Sum(nil))
}
