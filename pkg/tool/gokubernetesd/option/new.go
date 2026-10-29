package option

func New() *Server {
	return &Server{Clusters: BuildClusters}
}
