package authdto

type LoginResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(225)" json:"email" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Status   string `json:"status" form:"status"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Role     string `json:"role"`
}

type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Gender   string `gorm:"type: varchar(255)" json:"gender"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Address  string `gorm:"type: varchar(255)" json:"address"`
	Status   string `gorm:"type: varchar(50)"  json:"status"`
	Role     string `gorm:"type: varchar(100)" json:"role"`
}

type RegisterResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname" form:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(225)" form:"email" json:"email" validate:"required"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Status   string `json:"status" form:"status"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
