package request

func (q *Request) WithBody(
	contentType string,
	body []byte,
) *Request {
	q.ContentType = contentType
	q.Body = body

	return q
}
