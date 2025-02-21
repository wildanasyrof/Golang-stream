package dto

type CreateAnimeRequest struct {
	Title       string   `json:"title" validate:"required"`
	AltTitles   string   `json:"alt_titles"`
	Chapters    string   `json:"chapters"`
	Studio      string   `json:"studio"`
	Year        string   `json:"year"`
	Rating      float64  `json:"rating"`
	Synopsis    string   `json:"synopsis"`
	ImageSource string   `json:"image_source"`
	Categories  []string `json:"categories"` // Category names
}

type UpdateAnimeRequest struct {
	Title       string   `json:"title,omitempty"`
	AltTitles   string   `json:"alt_titles,omitempty"`
	Chapters    string   `json:"chapters,omitempty"`
	Studio      string   `json:"studio,omitempty"`
	Year        string   `json:"year,omitempty"`
	Rating      float64  `json:"rating,omitempty"`
	Synopsis    string   `json:"synopsis,omitempty"`
	ImageSource string   `json:"image_source,omitempty"`
	Categories  []string `json:"categories,omitempty"` // List of category names
}
