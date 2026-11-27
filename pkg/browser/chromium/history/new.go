package history

func New(currentIndex int64) *Result {
	return &Result{CurrentIndex: currentIndex}
}
