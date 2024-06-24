package schemas

import (
	m "github.com/adam-fraga/adhora/models"
	"gorm.io/gorm"
)

func CreatePagesTable(db *gorm.DB) {
	db.AutoMigrate(&m.Page{})
}

func DropPagesTable(db *gorm.DB) {
	db.Migrator().DropTable(&m.Page{})
}
