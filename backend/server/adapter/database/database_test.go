package database_test

/*import (
	"testing"

	"github.com/Mushus/trashbox/backend/server/adapter/database"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type rawUser struct {
	login    string
	password string
}

func TestVerifyUser(t *testing.T) {
	table := []struct {
		in  rawUser
		out bool
	}{
		{
			in: rawUser{
				login:    "admin",
				password: "admin",
			},
			out: true,
		},
		{
			in: rawUser{
				login:    "admin",
				password: "hoge",
			},
			out: false,
		},
		{
			in: rawUser{
				login:    "hoge",
				password: "hoge",
			},
			out: false,
		},
	}

	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("%v", err)
	}
	repository := database.ProvideUserDatastore(db)

	for _, row := range table {
		_, ok, err := repository.FindByLogin(row.in.login, row.in.password)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if ok != row.out {
			t.Fatalf("ok is %v, expect %v: user = %#v", ok, row.out, row.in)
		}
	}
}*/
