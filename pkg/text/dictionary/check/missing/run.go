package missing

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/text/dictionary"
	"github.com/funtimecoding/soil/pkg/text/dictionary/constant"
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

	fmt.Printf("Check %d words\n", total)
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
					if constant.Skip[d.Name()] {
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
	fmt.Printf("Scanned %d files\n", scanned)
	unused := make(map[string][]string)
	used := 0

	for _, u := range usage {
		if u.Used {
			used++
		} else {
			if unused[u.Category] == nil {
				unused[u.Category] = make([]string, 0)
			}

			unused[u.Category] = append(
				unused[u.Category],
				u.Word,
			)
		}
	}

	fmt.Printf("Results:\n")
	fmt.Printf(
		"Used: %d/%d words (%.1f%%)\n",
		used,
		total,
		float64(used)/float64(total)*100,
	)
	fmt.Printf("Unused: %d words\n", total-used)

	if len(unused) == 0 {
		fmt.Println("No unused dictionary words")

		return
	}

	fmt.Println("Unused words by category:")

	for c, words := range unused {
		fmt.Printf("\n# %s (%d unused):\n", c, len(words))

		for _, w := range words {
			fmt.Printf("  %s\n", w)
		}
	}
}
