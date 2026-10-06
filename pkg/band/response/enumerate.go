package response

type Enumerate struct {
	Context string `xml:"Body>EnumerateResponse>EnumerationContext"`
}
