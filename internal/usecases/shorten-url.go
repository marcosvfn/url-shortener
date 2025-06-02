package usecases

import (
	"github.com/marcosvfn/url-shortener/internal/domain/url"
)

type URLService struct {
	repo url.Repository
}

func NewURLService(repo url.Repository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) ShortenURL(originalURL string) (string, error) {
	return s.repo.(*url.Service).ShortenURL(originalURL)
}

func (s *URLService) GetOriginalURL(shortCode string) (string, error) {
	return s.repo.(*url.Service).GetOriginalURL(shortCode)
}
