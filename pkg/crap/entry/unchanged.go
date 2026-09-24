package entry

import "math"

func (e *Entry) Unchanged(tolerance float64) bool {
	return !e.IsNew() && math.Abs(e.Delta()) <= tolerance
}
