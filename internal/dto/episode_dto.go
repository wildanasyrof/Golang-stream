package dto

type CreateEpisodeRequest struct {
	EpisodeNumber int    `json:"episode_number" validate:"required,min=1"`
	Title         string `json:"title" validate:"required"`
	VideoURL      string `json:"video_url" validate:"required,url"`
}

type UpdateEpisodeRequest struct {
	EpisodeNumber int    `json:"episode_number,omitempty"`
	Title         string `json:"title,omitempty"`
	VideoURL      string `json:"video_url,omitempty"`
}
