package aliased

import st "example/pkg/outer/store"

func Run() *st.Store {
	return &st.Store{}
}
