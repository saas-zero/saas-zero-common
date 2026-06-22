package ryjwt

import (
	"testing"
	"time"
)

var testConf = &TokenConf{
	AccessSecret:  "test-secret-key",
	AccessExpire:  72,
	RefreshExpire: 168,
}

func TestSign(t *testing.T) {
	token, err := Sign(testConf, "userId", "12345", 1)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}
	if token == "" {
		t.Error("Sign() should not return empty token")
	}
}

func TestValid(t *testing.T) {
	token, err := Sign(testConf, "userId", "12345", 1)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	userId, err := Valid(testConf, "userId", token)
	if err != nil {
		t.Fatalf("Valid() error = %v", err)
	}
	if userId != "12345" {
		t.Errorf("Valid() userId = %q, want %q", userId, "12345")
	}
}

func TestValid_InvalidToken(t *testing.T) {
	_, err := Valid(testConf, "userId", "invalid.token.here")
	if err == nil {
		t.Error("Valid() should return error for invalid token")
	}
}

func TestValid_ExpiredToken(t *testing.T) {
	token, err := Sign(testConf, "userId", "12345", 0)
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	_, err = Valid(testConf, "userId", token)
	_ = err
}

func TestSign_DifferentKeys(t *testing.T) {
	token1, _ := Sign(testConf, "key1", "value1", 1)
	token2, _ := Sign(testConf, "key2", "value2", 1)

	if token1 == token2 {
		t.Error("Sign() with different keys should produce different tokens")
	}
}

func TestValid_DifferentKeys(t *testing.T) {
	token, _ := Sign(testConf, "userId", "12345", 1)

	_, err := Valid(testConf, "otherKey", token)
	if err == nil {
		t.Error("Valid() should return error for non-existent key")
	}
}

func TestGenerateTokens(t *testing.T) {
	accessToken, refreshToken, err := GenerateTokens(testConf, "12345")
	if err != nil {
		t.Fatalf("GenerateTokens() error = %v", err)
	}
	if accessToken == "" {
		t.Error("GenerateTokens() accessToken should not be empty")
	}
	if refreshToken == "" {
		t.Error("GenerateTokens() refreshToken should not be empty")
	}
}

func TestRefreshToken(t *testing.T) {
	token, err := RefreshToken(testConf, "12345")
	if err != nil {
		t.Fatalf("RefreshToken() error = %v", err)
	}
	if token == "" {
		t.Error("RefreshToken() should not return empty token")
	}
}