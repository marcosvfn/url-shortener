package url

import (
	"crypto/rand"
	"encoding/base64"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ShortenURL(originalURL string) (string, error) {
	shortCode := generateShortCode()
	url, err := NewURL(originalURL, shortCode)
	if err != nil {
		return "", err
	}

	err = s.repo.Save(url)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *Service) GetOriginalURL(shortCode string) (string, error) {
	url, err := s.repo.FindByShortCode(shortCode)
	if err != nil {
		return "", err
	}
	return url.OriginalURL, nil
}

func generateShortCode() string {
	bytes := make([]byte, 6)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)[:8]
}
