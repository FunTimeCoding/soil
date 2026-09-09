package pipeline

import (
	"golang.org/x/mod/semver"
	"log"
)

func LatestSemantic(v []*Pipeline) *Pipeline {
	if len(v) == 0 {
		log.Panic("empty slice")
	}

	var result *Pipeline

	for _, e := range v {
		if !semver.IsValid(e.Reference) {
			continue
		}

		if result == nil {
			result = e

			continue
		}

		if semver.Compare(e.Reference, result.Reference) > 0 {
			result = e
		}
	}

	return result
}
