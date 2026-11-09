package part

type Echo int

func (e Echo) Size() int {
	return int(e)
}
