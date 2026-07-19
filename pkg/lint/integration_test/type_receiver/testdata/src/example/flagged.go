package example

type Foo struct{} // want `multiple types with receivers in one package; extract to subpackage`

func (f *Foo) Method() {}
