package api_client

import "errors"

// InMemoryClient ApiClientのインメモリ実装. FakeDataを返す
type InMemoryClient struct {
	Store map[string]Data
}

func NewInMemoryClient() *InMemoryClient {
	client := &InMemoryClient{
		Store: make(map[string]Data),
	}

	// 初期データの設定
	client.Store["1"] = Data{ID: "1", Name: "Fake Data 1"}
	client.Store["2"] = Data{ID: "2", Name: "Fake Data 2"}
	client.Store["3"] = Data{ID: "3", Name: "Fake Data 3"}
	return client
}

func (c *InMemoryClient) FetchData() ([]Data, error) {
	if len(c.Store) == 0 {
		return nil, errors.New("no data available")
	}

	var data []Data
	for _, v := range c.Store {
		data = append(data, v)
	}
	return data, nil
}
