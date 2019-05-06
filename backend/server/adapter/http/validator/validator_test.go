package validator_test

import (
	"reflect"
	"testing"

	"github.com/Mushus/trashbox/backend/server/adapter/http/validator"
)

func TestValidator(t *testing.T) {
	v := validator.NewValidator()
	prm := validator.LoginParam{} // TODO: fix
	err := v.Validate(prm)
	if err == nil {
		t.Fatal("expect error")
	}

	result := validator.ReportValidation(err)
	want := validator.ValidationResult{
		"Login":    {"ログイン名 を入力してください"},
		"Password": {"パスワード を入力してください"},
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("result is %#v, expect %#v", result, want)
	}
}
