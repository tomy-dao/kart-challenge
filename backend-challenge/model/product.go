package model

type Product struct {
	ID string `json:"id"`
	Image Image `json:"image"`
	Name string `json:"name"`
	Category string `json:"category"`
	Price float64 `json:"price"`
}

type Image struct {
	Thumbnail string `json:"thumbnail"`
	Mobile string `json:"mobile"`
	Tablet string `json:"tablet"`
	Desktop string `json:"desktop"`
}
