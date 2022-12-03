package db

import (
	"fmt"
	"net/url"
	"strings"
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

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

func normalizeURL(urlString string) (string, error) {
	trimmed := strings.TrimSpace(urlString)
	parsed, err := url.Parse(trimmed)
	if err != nil {
		return "", err
	}

	if (parsed.Scheme != "http") && (parsed.Scheme != "https") {
		message := fmt.Sprintf("Invalid URL scheme: %s", parsed.Scheme)
		return "", &ValidationError{Field: "url", Message: message}
	}

	return trimmed, nil
}

func ShortenURL(url string) (*URL, error) {
	normalized, err := normalizeURL(url)
	if err != nil {
		return nil, err
	}

	record := &URL{Url: normalized}

	result := DB.Create(&record)

	if result.Error != nil {
		return nil, result.Error
	}

	return record, nil
}
