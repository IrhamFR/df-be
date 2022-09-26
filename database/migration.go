package database

import (
	"dumbflix/models"
	"dumbflix/pkg/mysql"
	"fmt"
)

// migration when running XAMPP
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Film{},
		&models.Episode{},
		&models.Transaction{},
		&models.Category{},
		&models.Profile{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
