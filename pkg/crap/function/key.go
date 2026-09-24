package function

func (f *Function) Key() string {
	return NewKey(f.Package, f.File, f.Line)
}
