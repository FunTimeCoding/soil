package update_request

func New(
	status string,
	assignedTo string,
) *Request {
	return &Request{Status: status, AssignedTo: assignedTo}
}
