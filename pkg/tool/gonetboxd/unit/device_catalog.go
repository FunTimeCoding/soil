package unit

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/mock_client"

func deviceCatalog() *mock_client.Client {
	result := mock_client.New()
	result.SeedDeviceRole("server")
	result.SeedDeviceType("r640")
	result.SeedSite("home")

	return result
}
