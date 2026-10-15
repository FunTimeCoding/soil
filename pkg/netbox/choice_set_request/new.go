package choice_set_request

func New(choices []string) *Request {
	result := &Request{Choices: [][]string{}}

	for _, v := range choices {
		result.Choices = append(result.Choices, []string{v, v})
	}

	return result
}
