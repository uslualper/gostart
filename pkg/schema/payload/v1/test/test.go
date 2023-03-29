package test

type Test struct {
	Message string `json:"message" validate:"required"`
}
