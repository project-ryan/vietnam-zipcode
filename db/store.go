package db

type Store interface {
	Querier
}

type JsonStore struct {
	*Queries
}

func NewStore(data map[string]Location) Store {
	return &JsonStore{
		Queries: New(data),
	}
}
