package schemas

import (
	m "github.com/adam-fraga/adhora/models"
	"gorm.io/gorm"
)

func CreateUsersTable(db *gorm.DB) {
	db.AutoMigrate(&m.User{})
}

func DropUsersTable(db *gorm.DB) {
	db.Migrator().DropTable(&m.User{})
}
