package dto

type FavoriteRequest struct {
	AnimeID uint `json:"anime_id" validate:"required"`
}
