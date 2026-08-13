package store

import "sort"

func ComputeFacets(
	results []SearchResult,
	threshold int,
) []Facet {
	counts := map[string]map[string]int{}

	for _, r := range results {
		for key, value := range r.Metadata {
			if counts[key] == nil {
				counts[key] = map[string]int{}
			}

			counts[key][value]++
		}
	}

	keys := make([]string, 0, len(counts))

	for k := range counts {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var result []Facet

	for _, key := range keys {
		values := counts[key]
		hasAggregation := false

		for _, count := range values {
			if count > 1 {
				hasAggregation = true

				break
			}
		}

		if !hasAggregation {
			continue
		}

		f := Facet{Key: key, Distinct: len(values)}

		if f.Distinct <= threshold {
			f.Values = values
		}

		result = append(result, f)
	}

	return result
}
