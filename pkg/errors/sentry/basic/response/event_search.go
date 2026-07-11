package response

type EventSearch struct {
	Rows []EventRow `json:"data"`
}
