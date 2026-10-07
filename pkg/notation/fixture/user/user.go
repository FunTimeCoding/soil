package user

type User struct {
	Name    string         `json:"name"`
	Unknown map[string]any `json:"-"`
}
