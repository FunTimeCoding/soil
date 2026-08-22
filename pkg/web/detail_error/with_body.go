package detail_error

func (e *Detail) WithBody(body []byte) *Detail {
	e.Body = body

	return e
}
