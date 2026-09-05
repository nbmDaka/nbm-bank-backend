package password

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)


func TestHashPassword(t *testing.T) {

	rawPassword := "secret123"


	hash, err := Hash(
		rawPassword,
	)


	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}


	if hash == rawPassword {

		t.Fatalf(
			"password was not hashed",
		)
	}


	err = bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(rawPassword),
	)


	if err != nil {

		t.Fatalf(
			"hash does not match password: %v",
			err,
		)

	}
}