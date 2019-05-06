package user_test

import (
	"testing"

	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/Mushus/trashbox/backend/server/app/user"
)

func TestUserPassword(t *testing.T) {
	u, _ := user.NewUser(property.User{Login: "user", Password: "password"})

	if u.VerifyPassword("hello") {
		t.Fatalf("expect: incorrect, actual: correct")
	}
	if !u.VerifyPassword("password") {
		t.Fatalf("expect: correct, actual: incorrect")
	}
}
