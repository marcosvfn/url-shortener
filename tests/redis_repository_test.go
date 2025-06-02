package tests

import (
	"errors"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/marcosvfn/url-shortener/internal/domain/url"
	"github.com/marcosvfn/url-shortener/internal/infrastructure/redis"
)

func TestRedisRepository_Save(t *testing.T) {
	db, mock := redismock.NewClientMock()
	repo := redis.NewRedisRepositoryForTest(db)

	urlEntity, err := url.NewURL("https://example.com", "abc123")
	if err != nil {
		t.Fatalf("Failed to create URL: %v", err)
	}

	mock.ExpectSet("abc123", "https://example.com", 0).SetVal("OK")

	err = repo.Save(urlEntity)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestRedisRepository_Save_Error(t *testing.T) {
	db, mock := redismock.NewClientMock()
	repo := redis.NewRedisRepositoryForTest(db)

	urlEntity, err := url.NewURL("https://example.com", "abc123")
	if err != nil {
		t.Fatalf("Failed to create URL: %v", err)
	}

	expectedErr := errors.New("redis error")
	mock.ExpectSet("abc123", "https://example.com", 0).SetErr(expectedErr)

	err = repo.Save(urlEntity)
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestRedisRepository_FindByShortCode_Success(t *testing.T) {
	db, mock := redismock.NewClientMock()
	repo := redis.NewRedisRepositoryForTest(db)

	mock.ExpectGet("abc123").SetVal("https://example.com")

	result, err := repo.FindByShortCode("abc123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected a URL, got nil")
	}
	if result.OriginalURL != "https://example.com" {
		t.Errorf("Expected OriginalURL to be %s, got %s", "https://example.com", result.OriginalURL)
	}
	if result.ShortCode != "abc123" {
		t.Errorf("Expected ShortCode to be %s, got %s", "abc123", result.ShortCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestRedisRepository_FindByShortCode_NotFound(t *testing.T) {
	db, mock := redismock.NewClientMock()
	repo := redis.NewRedisRepositoryForTest(db)

	mock.ExpectGet("abc123").RedisNil()

	result, err := repo.FindByShortCode("abc123")
	if err == nil || err.Error() != "URL not found" {
		t.Errorf("Expected error 'URL not found', got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestRedisRepository_FindByShortCode_Error(t *testing.T) {
	db, mock := redismock.NewClientMock()
	repo := redis.NewRedisRepositoryForTest(db)

	expectedErr := errors.New("redis error")
	mock.ExpectGet("abc123").SetErr(expectedErr)

	result, err := repo.FindByShortCode("abc123")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
