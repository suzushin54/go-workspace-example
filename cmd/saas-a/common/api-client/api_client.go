package api_client

type ApiClient interface {
	FetchData() ([]Data, error)
}

type Data struct {
	ID   string
	Name string
}
