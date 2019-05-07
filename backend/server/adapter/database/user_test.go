package database

import "testing"

func TestUserToObject(t *testing.T) {
	password := "password"
	ud := userData{
		Password: "$2a$10$7TxKqk3wRYKMO8sEgF6yA.BQ8SDp44cBZKapAcGb.dmHI2Zcnw8V6",
	}

	u, err := ud.toObject()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	expect := true
	actual := u.VerifyPassword(password)
	if expect != actual {
		t.Fatalf("expect: %v, actual: %v", expect, actual)
	}
}
