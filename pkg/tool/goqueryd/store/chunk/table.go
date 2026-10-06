package chunk

type table struct {
	start     int
	headerEnd int
	end       int
	rows      []line
	heading   string
	caption   string
}
