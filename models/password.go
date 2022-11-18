package models

type ForgotInitial struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type ForgotValidation struct {
	ForgotKey string `json:"forgot_key" validate:"required"`
	Key       string `json:"key" validate:"required"`
	Password  string `json:"password"  validate:"required" gorm:"password"`
}
