package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "my_secure_password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}
	if hash == "" {
		t.Fatal("expected non-empty hash")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "my_secure_password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	match, err := CheckPasswordHash(password, hash)
	if err != nil {
		t.Fatalf("CheckPasswordHash failed: %v", err)
	}
	if !match {
		t.Fatal("expected password to match hash")
	}

	match, err = CheckPasswordHash("wrong_password", hash)
	if err != nil {
		t.Fatalf("CheckPasswordHash failed: %v", err)
	}
	if match {
		t.Fatal("expected wrong password to not match hash")
	}
}
