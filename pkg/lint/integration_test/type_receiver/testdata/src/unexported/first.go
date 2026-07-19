package unexported

type handler struct{} // want `multiple types with receivers in one package; extract to subpackage`

func (h *handler) serve() {}
