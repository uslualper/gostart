package validate

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
