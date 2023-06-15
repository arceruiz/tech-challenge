package entities

type Product struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Price     float64 `json:"price"`
	Category  string  `json:"category"`
	ImagePath string  `json:"imagePath"`
}
