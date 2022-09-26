package authdto

type RegisterRequest struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Gender   string `json:"gender" gorm:"type:varchar(50)"`
	Phone    string `json:"phone" gorm:"type:varchar(50)"`
	Address  string `json:"address" gorm:"type:varchar(225)"`
	Status   string `json:"status" gorm:"type: varchar(50)"`
	Role     string `gorm:"type: varchar(255)" json:"role"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
