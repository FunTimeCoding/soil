package mock_client

type Client struct {
	Failure  error
	Uploaded []string
	Closed   int
}
