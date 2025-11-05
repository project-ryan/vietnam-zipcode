package db

type Queries struct {
	data map[string]Location
}

func New(data map[string]Location) *Queries {
	return &Queries{
		data: data,
	}
}
