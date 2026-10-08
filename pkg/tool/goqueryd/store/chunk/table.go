package chunk

type Table struct {
	start     int
	headerEnd int
	end       int
	rows      []Line
	heading   string
	caption   string
}
