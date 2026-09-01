package missing

import (
	"github.com/funtimecoding/soil/pkg/console"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	text "github.com/funtimecoding/soil/pkg/text/constant"
	"github.com/funtimecoding/soil/pkg/text/dictionary"
	"io/fs"
	"path/filepath"
	"strings"
)

func Run() {
	usage := make(map[string]*dictionary.WordUsage)
	total := 0

	for _, c := range dictionary.Read(dictionary.ResolvePath()) {
		for _, w := range c.Words {
			usage[strings.ToLower(w)] = dictionary.NewWordUsage(
				w,
				c.Name,
				false,
			)
			total++
		}
	}

	console.Format("Check %d words\n", total)
	scanned := 0
	errors.PanicOnError(
		filepath.WalkDir(
			library.CurrentDirectory,
			func(
				path string,
				d fs.DirEntry,
				e error,
			) error {
				if e != nil {
					return e
				}

				if d.IsDir() {
					if text.DictionarySkip[d.Name()] {
						return filepath.SkipDir
					}

					return nil
				}

				if !dictionary.IncludeFile(d.Name()) {
					return nil
				}

				dictionary.ScanFile(path, usage)
				scanned++

				return nil
			},
		),
	)
	console.Format("Scanned %d files\n", scanned)
	unused := make(map[string][]string)
	used := 0

	for _, u := range usage {
		if u.Used {
			used++
		} else {
			if unused[u.Category] == nil {
				unused[u.Category] = make([]string, 0)
			}

			unused[u.Category] = append(unused[u.Category], u.Word)
		}
	}

	console.Format("Results:\n")
	console.Format(
		"Used: %d/%d words (%.1f%%)\n",
		used,
		total,
		float64(used)/float64(total)*100,
	)
	console.Format("Unused: %d words\n", total-used)

	if len(unused) == 0 {
		console.Line("No unused dictionary words")

		return
	}

	console.Line("Unused words by category:")

	for c, words := range unused {
		console.Format("\n# %s (%d unused):\n", c, len(words))

		for _, w := range words {
			console.Format("  %s\n", w)
		}
	}
}
