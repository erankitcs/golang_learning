package module1

// Budget stores budget information
type Budget struct {
	Max   float32
	Items []Item
}

// Item stores item information
type Item struct {
	Description string
	Price       float32
}
