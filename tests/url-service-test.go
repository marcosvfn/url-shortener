package tests

import (
	"errors"
	"testing"

	"github.com/marcosvfn/url-shortener/internal/domain/url"
)

type mockURLRepository struct {
	urls map[string]string
}

func (m *mockURLRepository) Save(url *url.URL) error {
	m.urls[url.ShortCode] = url.OriginalURL
	return nil
}

func (m *mockURLRepository) FindByShortCode(shortCode string) (*url.URL, error) {
	originalURL, exists := m.urls[shortCode]
	if !exists {
		return nil, errors.New("URL not found")
	}
	return &url.URL{OriginalURL: originalURL, ShortCode: shortCode}, nil
}

func TestURLService_ShortenURL(t *testing.T) {
	repo := &mockURLRepository{urls: make(map[string]string)}
	service := url.NewService(repo)

	originalURL := "https://example.com"
	shortCode, err := service.ShortenURL(originalURL)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(shortCode) != 8 {
		t.Errorf("Expected short code length 8, got %d", len(shortCode))
	}
}

func TestURLService_GetOriginalURL(t *testing.T) {
	repo := &mockURLRepository{urls: make(map[string]string)}
	service := url.NewService(repo)

	originalURL := "https://example.com"
	shortCode, _ := service.ShortenURL(originalURL)

	retrievedURL, err := service.GetOriginalURL(shortCode)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if retrievedURL != originalURL {
		t.Errorf("Expected %s, got %s", originalURL, retrievedURL)
	}
}
