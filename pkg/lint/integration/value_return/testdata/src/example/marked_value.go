package example

import "example/inner"

func MarkedByValue() inner.Marked {
	return inner.Marked{Value: 1}
}

func BareDirectiveByValue() inner.BareDirective {
	return inner.BareDirective{Value: 1}
}

func UnmarkedByValue() inner.Unmarked {
	return inner.Unmarked{Value: 1}
}
