package environment

func (e *Environment) Load(path string) error {
	overlay, loadError := eval(path, e.base)

	if loadError != nil {
		return loadError
	}

	current, exportError := exported(path, e.base)

	if exportError != nil {
		return exportError
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	for key := range e.exported {
		if _, present := current[key]; !present {
			e.deleted[key] = struct{}{}
		}
	}

	for key := range current {
		delete(e.deleted, key)
	}

	e.overlay = overlay
	e.exported = current

	return nil
}
