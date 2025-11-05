package db

// Location represents a geographic location with province and ward
type Location struct {
	Province string `json:"province"`
	Ward     string `json:"ward"`
}

type Data map[string]Location
