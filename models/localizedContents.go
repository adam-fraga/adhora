package models

import (
	d "github.com/adam-fraga/adhora/db"
	"gorm.io/gorm"
)

type LocalizedContent struct {
	gorm.Model
	ID        int    `json:"id"`
	FrTitle   string `json:"titlefr"`
	EnTitle   string `json:"titleen"`
	FrContent string `json:"frcontent"`
	EnContent string `json:"encontent"`
	Page      Page   `json:"page"`
}

// GetLocalizedContentByPageName returns a localized content by page name
func (lc *LocalizedContent) GetLocalizedContentByPageName(page Page, db *d.DB) LocalizedContent {
	db.NewConnection()
	defer db.Close()

	db.Client.AutoMigrate(&LocalizedContent{})
	var localizedContent LocalizedContent
	db.Client.Where("page_name = ?", page.Name).First(&localizedContent)
	return localizedContent
}

// GetLocalizedContent returns a localized content by page
func (lc *LocalizedContent) GetLocalizedContent(page Page, db *d.DB) LocalizedContent {
	db.NewConnection()
	defer db.Close()

	db.Client.AutoMigrate(&LocalizedContent{})
	var localizedContent LocalizedContent
	db.Client.First(&localizedContent, "page_id = ?", page.ID)
	return localizedContent
}

// GetLocalizedContents returns all localized contents
func (lc *LocalizedContent) GetLocalizedContents(db *d.DB) []LocalizedContent {

	db.NewConnection()
	defer db.Close()

	db.Client.AutoMigrate(&LocalizedContent{})
	var localizedContents []LocalizedContent
	db.Client.Find(&localizedContents)
	return localizedContents
}

// SetLocalizedContent sets a localized SetLocalizedContent()
func (lc *LocalizedContent) SetLocalizedContent(localizedContent LocalizedContent, db *d.DB) {

	db.NewConnection()
	defer db.Close()

	db.Client.Save(&localizedContent)
}

func (lc *LocalizedContent) DeleteLocalizedContent(db *d.DB) {

	db.NewConnection()
	defer db.Close()

	db.Client.Delete(&LocalizedContent{}, lc.ID)
}
