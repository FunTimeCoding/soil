package gofix

type FileEdit struct {
	offset  int
	length  int
	newText string
}
