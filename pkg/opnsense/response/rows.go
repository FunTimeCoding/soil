package response

type Rows[T any] struct {
	Rows []T `json:"rows"`
}
