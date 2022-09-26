package filmsdto

type FilmResponse struct {
	ID               int    `json:"id" gorm:"primary_key:auto_increment"`
	Title            string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm    string `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Description      string `json:"desc" gorm:"type:text" form:"description"`
	Year             int    `json:"year" form:"year" gorm:"type: int"`
	TitleEpisode     string `json:"titleEpisode" form:"titleEpisode" gorm:"type: varchar(255)"`
	ThumbnailEpisode string `json:"thumbnailEpisode" form:"thumbnailEpisode" gorm:"type: varchar(255)"`
	LinkFilm         string `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`
}
