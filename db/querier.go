package db

// Querier defines the interface for zipcode operations
type Querier interface {
	GetLocation(zipcode string) (*Location, error)
}
