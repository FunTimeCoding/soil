package token_summary

import "slices"

func NewSpread(values []int) *Spread {
	if len(values) == 0 {
		return &Spread{}
	}

	sorted := slices.Clone(values)
	slices.Sort(sorted)
	total := 0

	for _, v := range sorted {
		total += v
	}

	return &Spread{
		Median:      at(sorted, 50),
		Ninetieth:   at(sorted, 90),
		NinetyNinth: at(sorted, 99),
		Maximum:     sorted[len(sorted)-1],
		Total:       total,
	}
}
