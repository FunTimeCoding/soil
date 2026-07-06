package gofix

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"path/filepath"
	"strings"
)

func runSurvey(patterns []string) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	counts := make(map[string]int)
	examples := make(map[string][]string)
	all, _ := load("", patterns)

	for _, p := range all {
		generated := buildGeneratedSet(p)

		for i, o := range p.TypesInfo.Defs {
			if o == nil {
				continue
			}

			f := p.Fset.File(i.Pos()).Name()

			if generated[f] {
				continue
			}

			if strings.HasSuffix(filepath.Base(f), constant.TestSuffix) {
				continue
			}

			for _, s := range segment.Segments(i.Name) {
				if whitelist[s] {
					continue
				}

				counts[s]++

				if len(examples[s]) < 3 {
					examples[s] = append(examples[s], i.Name)
				}
			}
		}
	}

	printSurvey(counts, examples)
}
