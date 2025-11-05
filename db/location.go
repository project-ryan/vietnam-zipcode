package db

import "fmt"

func (q *Queries) GetLocation(zipcode string) (*Location, error) {
	location, ok := q.data[zipcode]
	fmt.Println("get ", zipcode, "->", location, "found:", ok)

	if !ok {
		return nil, nil
	}

	return &location, nil
}
