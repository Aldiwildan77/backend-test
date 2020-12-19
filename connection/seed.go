package connection

import (
	"backend-majoo-test/entities"
	"log"

	"gorm.io/gorm"
)

var users = []entities.User{
	entities.User{
		Username: "aldiwildan",
		Password: "123123",
		FullName: "Muhammad Wildan",
	},
	entities.User{
		Username: "wildansyah",
		Password: "123123",
		FullName: "Wildan Aldiansyah",
	},
}

// Seed start to seed table
func Seed(db *gorm.DB) {
	if err := db.Debug().Migrator().DropTable(&entities.User{}); err != nil {
		log.Fatalf("failed to drop table: %v", err)
	}

	if err := db.Debug().AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("failed to migrate table: %v", err)
	}

	for _, user := range users {
		err := db.Debug().Create(&user).Error
		if err != nil {
			log.Fatalf("failed to create user: %v", err)
		}
	}
}
