package models

type Film struct {
	ID            int              `json:"id" gorm:"primary_key:auto_increment"`
	Title         string           `json:"title" gorm:"type: varchar(255)" form:"title"`
	ThumbnailFilm string           `json:"thumbnail" form:"image" gorm:"type:varchar(225)"`
	Year          string           `json:"year" gorm:"type:varchar(50)" form:"year"`
	CategoryID    int              `json:"category_id" gorm:"type:int" form:"category_id"`
	Category      CategoryResponse `json:"category"`
	Description   string           `json:"description"  gorm:"type: text" form:"description"`
	TitleEpisode  string           `json:"titleEpisode" gorm:"type: varchar(255)" form:"titleEpisode"`
	// ThumbnailEpisode string           `json:"thumbnailEpisode" form:"image" gorm:"type: varchar(255)"`
	LinkFilm string `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`
}

type FilmResponse struct {
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	ThumbnailFilm string           `json:"thumbnail"`
	Year          string           `json:"year"`
	CategoryID    int              `json:"category_id"`
	Category      CategoryResponse `json:"category"`
	Description   string           `json:"description"`
	TitleEpisode  string           `json:"titleEpisode"`
	// ThumbnailEpisode string           `json:"thumbnailEpisode"`
	LinkFilm string `json:"linkfilm"`
}

type FilmCategoryResponse struct {
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	ThumbnailFilm string           `json:"thumbnail"`
	Year          string           `json:"year"`
	CategoryID    int              `json:"category_id"`
	Category      CategoryResponse `json:"category"`
	Description   string           `json:"description"`
	TitleEpisode  string           `json:"titleEpisode"`
	// ThumbnailEpisode string           `json:"thumbnailEpisode"`
	LinkFilm string `json:"linkfilm"`
}

type FilmEpisodeResponse struct {
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	ThumbnailFilm string           `json:"thumbnail"`
	Year          string           `json:"year"`
	CategoryID    int              `json:"category_id"`
	Category      CategoryResponse `json:"category"`
	TitleEpisode  string           `json:"titleEpisode"`
	// ThumbnailEpisode string           `json:"thumbnailEpisode"`
	LinkFilm string `json:"linkfilm"`
}

func (FilmCategoryResponse) TableName() string {
	return "films"
}

func (FilmEpisodeResponse) TableName() string {
	return "films"
}

func (FilmResponse) TableName() string {
	return "films"
}
