package usecases

import (
	"github.com/marcosvfn/url-shortener/internal/domain/url"
)

type URLService struct {
	service *url.Service
}

func NewURLService(service *url.Service) *URLService {
	return &URLService{service: service}
}

func (s *URLService) ShortenURL(originalURL string) (string, error) {
	return s.service.ShortenURL(originalURL)
}

func (s *URLService) GetOriginalURL(shortCode string) (string, error) {
	return s.service.GetOriginalURL(shortCode)
}
