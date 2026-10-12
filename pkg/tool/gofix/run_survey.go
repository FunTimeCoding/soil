package gofix

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"github.com/funtimecoding/soil/pkg/tool/gofix/constant"
	"path/filepath"
	"strings"
)

func runSurvey(
	directory string,
	patterns []string,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	counts := make(map[string]int)
	examples := make(map[string][]string)
	all, _ := Load(directory, patterns)

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

			if strings.HasSuffix(filepath.Base(f), soil.TestSuffix) {
				continue
			}

			for _, s := range segment.Segments(i.Name) {
				if constant.Whitelist[s] {
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
