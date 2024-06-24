package models

import (
	d "github.com/adam-fraga/adhora/db"
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetPages returns all pages
func (p *Page) GetPage(db *d.DB) Page {

	db.NewConnection()
	defer db.Close()

	db.Client.AutoMigrate(&Page{})
	var page Page
	db.Client.First(&page, p.ID)
	return page
}

// GetPagesByName returns a page by name
func (p *Page) GetPageByName(db *d.DB) Page {

	db.NewConnection()
	defer db.Close()

	db.Client.AutoMigrate(&Page{})
	var page Page
	db.Client.Where("name = ?", p.Name).First(&page)
	return page
}

// SetPage sets a page
func (p *Page) SetPage(page Page, db *d.DB) {

	db.NewConnection()
	defer db.Close()

	db.Client.Save(&page)
}

// DeletePage deletes a page
func (p *Page) DeletePage(db *d.DB) {

	db.NewConnection()
	defer db.Close()

	db.Client.Delete(&Page{}, p.ID)
}
