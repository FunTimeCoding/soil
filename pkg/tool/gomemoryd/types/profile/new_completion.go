package profile

func NewCompletion(
	sessionName string,
	body string,
) *Completion {
	return &Completion{SessionName: sessionName, Body: body}
}
