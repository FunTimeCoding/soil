package duplicate

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/text/dictionary"
)

func Run() {
	categories := dictionary.Read(dictionary.ResolvePath())
	occurrences := make(map[string][]dictionary.WordLocation)
	total := 0

	for _, c := range categories {
		for _, w := range c.Words {
			occurrences[w] = append(
				occurrences[w],
				dictionary.WordLocation{Category: c.Name, Word: w},
			)
			total++
		}
	}

	duplicates := make(map[string][]dictionary.WordLocation)
	count := 0

	for w, locations := range occurrences {
		if len(locations) > 1 {
			duplicates[w] = locations
			count += len(locations) - 1
		}
	}

	fmt.Printf("Checked %d words\n", total)

	if len(duplicates) == 0 {
		fmt.Println("No duplicate words found")

		return
	}

	fmt.Printf(
		"Found %d duplicate words (%d extra occurrences):\n",
		len(duplicates),
		count,
	)

	for w, locations := range duplicates {
		fmt.Printf("\n'%s' appears %d times:\n", w, len(locations))

		for _, l := range locations {
			fmt.Printf("- in category: %s\n", l.Category)
		}
	}
}
