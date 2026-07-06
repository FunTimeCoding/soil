package face

type ProcessRegistry interface {
	Add(pid int)
	Remove(pid int)
	Contains(pid int) bool
}
