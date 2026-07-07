package jwt

import (
	"testing"
	"time"
)

const testSecret = "test-secret-key-for-unit-testing"

func TestSignAndParse_Success(t *testing.T) {
	claims := &Claims{
		UserId:    1001,
		TenantId:  2001,
		UserName:  "testuser",
		RoleCodes: []string{"admin", "user"},
	}
	token, err := Sign(testSecret, claims, time.Hour)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	parsed, err := Parse(token, testSecret)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if parsed.UserId != 1001 {
		t.Fatalf("expected UserId=1001, got %d", parsed.UserId)
	}
	if parsed.TenantId != 2001 {
		t.Fatalf("expected TenantId=2001, got %d", parsed.TenantId)
	}
	if parsed.UserName != "testuser" {
		t.Fatalf("expected UserName=testuser, got %s", parsed.UserName)
	}
	if len(parsed.RoleCodes) != 2 || parsed.RoleCodes[0] != "admin" || parsed.RoleCodes[1] != "user" {
		t.Fatalf("unexpected RoleCodes: %v", parsed.RoleCodes)
	}
}

func TestParse_WrongSecret(t *testing.T) {
	claims := &Claims{
		UserId:   1,
		TenantId: 1,
		UserName: "test",
	}
	token, err := Sign(testSecret, claims, time.Hour)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	_, err = Parse(token, "wrong-secret")
	if err == nil {
		t.Fatal("expected error when parsing with wrong secret")
	}
}

func TestParse_ExpiredToken(t *testing.T) {
	claims := &Claims{
		UserId:   1,
		TenantId: 1,
		UserName: "test",
	}
	token, err := Sign(testSecret, claims, -time.Hour)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	_, err = Parse(token, testSecret)
	if err == nil {
		t.Fatal("expected error when parsing expired token")
	}
}

func TestParse_MalformedToken(t *testing.T) {
	_, err := Parse("not-a-valid-token", testSecret)
	if err == nil {
		t.Fatal("expected error when parsing malformed token")
	}
}

func TestParse_EmptyToken(t *testing.T) {
	_, err := Parse("", testSecret)
	if err == nil {
		t.Fatal("expected error when parsing empty token")
	}
}

func TestSignAndParse_RoleCodesNil(t *testing.T) {
	claims := &Claims{
		UserId:   42,
		TenantId: 7,
		UserName: "nobody",
	}
	token, err := Sign(testSecret, claims, time.Hour)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	parsed, err := Parse(token, testSecret)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if parsed.UserId != 42 {
		t.Fatalf("expected UserId=42, got %d", parsed.UserId)
	}
}

func TestSignAndParse_EmptyRoleCodes(t *testing.T) {
	claims := &Claims{
		UserId:    99,
		TenantId:  99,
		UserName:  "roleless",
		RoleCodes: []string{},
	}
	token, err := Sign(testSecret, claims, time.Hour)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	parsed, err := Parse(token, testSecret)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(parsed.RoleCodes) != 0 {
		t.Fatalf("expected empty RoleCodes, got %v", parsed.RoleCodes)
	}
}

func TestParse_TokenNotExpired(t *testing.T) {
	claims := &Claims{
		UserId:   1,
		TenantId: 1,
		UserName: "fresh",
	}
	token, err := Sign(testSecret, claims, time.Hour)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	parsed, err := Parse(token, testSecret)
	if err != nil {
		t.Fatalf("Parse failed for valid token: %v", err)
	}
	_ = parsed
}

func TestSignAndParse_LongRoleCodes(t *testing.T) {
	roles := make([]string, 50)
	for i := range roles {
		roles[i] = "role-" + string(rune('A'+i%26))
	}
	claims := &Claims{
		UserId:    1,
		TenantId:  1,
		UserName:  "multirole",
		RoleCodes: roles,
	}
	token, err := Sign(testSecret, claims, time.Hour)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	parsed, err := Parse(token, testSecret)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(parsed.RoleCodes) != len(roles) {
		t.Fatalf("expected %d roles, got %d", len(roles), len(parsed.RoleCodes))
	}
}
