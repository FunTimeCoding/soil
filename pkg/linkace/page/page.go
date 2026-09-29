package page

type Page[T any] struct {
	Items       []T
	Total       int
	CurrentPage int
	LastPage    int
	PerPage     int
}
