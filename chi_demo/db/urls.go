package db

import (
	"time"
)

type URL struct {
	ID         uint      `gorm:"id" json:"id"`
	Url        string    `gorm:"url" json:"url"`
	ShortId    string    `gorm:"short_id" json:"short_id"`
	InsertedAt time.Time `gorm:"inserted_at" json:"inserted_at"`
	UpdatedAt  time.Time `gorm:"updated_at" json:"updated_at"`
}

func ListURLs() []URL {
	var urls []URL
	_ = *DB.Find(&urls)
	return urls
}
