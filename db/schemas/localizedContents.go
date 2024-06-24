package schemas

import (
	m "github.com/adam-fraga/adhora/models"
	"gorm.io/gorm"
)

func CreateLocalizedContentsTable(db *gorm.DB) {
	db.AutoMigrate(&m.LocalizedContent{})
}

func DropLocalizedContentsTable(db *gorm.DB) {
	db.Migrator().DropTable(&m.LocalizedContent{})
}
