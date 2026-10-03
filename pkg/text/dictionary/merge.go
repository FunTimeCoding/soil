package dictionary

import "sort"

func Merge(
	target string,
	sources ...string,
) int {
	categories := Read(target)
	index := make(map[string]int, len(categories))

	for i, c := range categories {
		index[c.Name] = i
	}

	added := 0

	for _, source := range sources {
		for _, sc := range Read(source) {
			if i, okay := index[sc.Name]; okay {
				existing := wordSet(categories[i].Words)

				for _, w := range sc.Words {
					if !existing[w] {
						categories[i].Words = append(categories[i].Words, w)
						existing[w] = true
						added++
					}
				}
			} else {
				index[sc.Name] = len(categories)
				categories = append(categories, sc)
				added += len(sc.Words)
			}
		}
	}

	for _, c := range categories {
		sort.Strings(c.Words)
	}

	Write(target, categories)

	return added
}
