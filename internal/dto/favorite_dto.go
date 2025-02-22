package dto

type AddFavoriteRequest struct {
	AnimeID uint `json:"anime_id" validate:"required"`
}

type RemoveFavoriteRequest struct {
	AnimeID uint `json:"anime_id" validate:"required"`
}
