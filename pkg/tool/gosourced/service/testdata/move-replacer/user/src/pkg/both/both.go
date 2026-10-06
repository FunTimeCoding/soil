package both

import "other.test/lib/source"

func Both() string {
	return source.Helper() + source.Keep()
}
