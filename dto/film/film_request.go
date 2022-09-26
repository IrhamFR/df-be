package filmsdto

type CreateFilmRequest struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)" validate:"required"`
	ThumbnailFilm string `json:"thumbnail" form:"image" gorm:"type: varchar(50)" validate:"required"`
	Year          string `json:"year" form:"year" gorm:"type: varchar(50)" validate:"required"`
	CategoryID    int    `json:"category_id" form:"category_id" validate:"required"`
	Description   string `json:"description" form:"description" validate:"required"`
	TitleEpisode  string `json:"titleEpisode" form:"titleEpisode" validate:"required"`
	// ThumbnailEpisode string `json:"thumbnailEpisode" form:"thumbnailEpisode" validate:"required"`
	LinkFilm string `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`
}

type UpdateFilmRequest struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)" validate:"required"`
	ThumbnailFilm string `json:"thumbnail" form:"image" gorm:"type: varchar(50)" validate:"required"`
	Year          string `json:"year" form:"year" gorm:"type: varchar(50)"`
	CategoryID    int    `json:"category_id" form:"category_id"`
	Description   string `json:"description" form:"description"`
	TitleEpisode  string `json:"titleEpisode" form:"titleEpisode" validate:"required"`
	// ThumbnailEpisode string `json:"thumbnailEpisode" form:"thumbnailEpisode" validate:"required"`
	LinkFilm string `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`
}
