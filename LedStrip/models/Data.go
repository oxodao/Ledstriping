package models

type Favorite struct {
	ID         string
	App        string
	Name       string
	Color      string
	Brightness uint64
	Speed      uint64
	Mode       uint64
}

type Data struct {
	Modes     []string
	Favorites []Favorite
}
