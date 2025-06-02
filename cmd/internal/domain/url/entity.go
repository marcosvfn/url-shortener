package url

import (
	"errors"
	"net/url"
)

type URL struct {
	OriginalURL string
	ShortCode   string
}

func NewURL(originalURL, shortCode string) (*URL, error) {
	if _, err := url.ParseRequestURI(originalURL); err != nil {
		return nil, errors.New("invalid URL")
	}
	return &URL{OriginalURL: originalURL, ShortCode: shortCode}, nil
}
