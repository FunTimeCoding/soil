package completion

func New(
	sessionName string,
	body string,
) *Completion {
	return &Completion{SessionName: sessionName, Body: body}
}
