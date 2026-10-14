package mock_closer

func New(result error) *Closer {
	return &Closer{result: result}
}
