package models

import (
	"errors"
	"pvg/databases"
	"time"
)

type ConfirmEmail struct {
	EmailKey string `json:"email_key" validate:"required"`
	Key      string `json:"key" validate:"required"`
}

type UserCreate struct {
	Id             string    `json:"id" gorm:"id"`
	Username       string    `json:"username" validate:"required" gorm:"username"`
	Firstname      string    `json:"firstname" validate:"required" gorm:"firstname"`
	Lastname       string    `json:"lastname"  validate:"required" gorm:"lastname"`
	Password       string    `json:"password"  validate:"required" gorm:"password"`
	Phone          string    `json:"phone" validate:"required" gorm:"phone"`
	Email          string    `json:"email" validate:"required" gorm:"email"`
	Birthday       time.Time `json:"birthday" validate:"required" gorm:"birthday"`
	EmailConfirmed int       `json:"email_confirmed" gorm:"email_confirmed"`
}

type User struct {
	Id             string    `json:"id" gorm:"id"`
	Username       string    `json:"username" gorm:"username"`
	Firstname      string    `json:"firstname" gorm:"firstname"`
	Lastname       string    `json:"lastname"  gorm:"lastname"`
	Password       string    `json:"password" gorm:"password"`
	Phone          string    `json:"phone" gorm:"phone"`
	Email          string    `json:"email" gorm:"email"`
	Birthday       time.Time `json:"birthday" gorm:"birthday"`
	EmailConfirmed int       `json:"email_confirmed" gorm:"email_confirmed"`
}

func (User) TableName() string {
	return "users"
}

func (UserCreate) TableName() string {
	return "users"
}

func (t *User) GetAllUser() (out *[]User, err error) {

	DB := databases.Connect()

	tx := DB.Table(t.TableName())
	tx.Find(&out)

	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	if tx.Error != nil {
		return nil, errors.New(tx.Error.Error())
	}

	return
}

func (t *User) GetOneUser() (out *User, err error) {

	DB := databases.Connect()
	tx := DB.Table(t.TableName())
	tx.First(&out, "id = ?", t.Id)

	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	if tx.Error != nil {
		return nil, errors.New(tx.Error.Error())
	}
	return
}

func (t *UserCreate) isUserExist() bool {

	DB := databases.Connect()
	var user User
	tx := DB.Table(t.TableName())
	tx.First(&user, "username = ? or email = ? or phone = ?", t.Username, t.Email, t.Phone)
	if tx.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}

func (t *User) isUpdateUserExist() bool {

	DB := databases.Connect()
	var user User
	tx := DB.Table(t.TableName())
	tx.First(&user, "username = ? or email = ? or phone = ?", t.Username, t.Email, t.Phone)
	if tx.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}

func (t *UserCreate) CreateUser() (out *UserCreate, err error) {

	isExist := t.isUserExist()
	if isExist {
		return nil, errors.New("user already exist")
	}

	var user UserCreate

	user.Id = t.Id

	// isExist := activity.isActivityExist(DB)
	// if !isExist {
	// 	return nil, errors.New("activity group doesn't exist")
	// }
	DB := databases.Connect()

	tx := DB.Table(t.TableName()).Create(&UserCreate{
		Username:       t.Username,
		Firstname:      t.Firstname,
		Lastname:       t.Lastname,
		Password:       t.Password,
		Phone:          t.Phone,
		Email:          t.Email,
		Birthday:       t.Birthday,
		EmailConfirmed: 0,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("create failed")
	}

	return
}

func (t *User) UpdateUser() (out *User, err error) {

	isExist := t.isUpdateUserExist()
	if isExist {
		return nil, errors.New("user already exist")
	}
	DB := databases.Connect()

	tx := DB.Table(t.TableName()).Where("id = ?", t.Id).Updates(&User{
		Username:       t.Username,
		Firstname:      t.Firstname,
		Lastname:       t.Lastname,
		Password:       t.Password,
		Phone:          t.Phone,
		Email:          t.Email,
		Birthday:       t.Birthday,
		EmailConfirmed: t.EmailConfirmed,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (t *User) ConfirmEmail() (err error) {

	DB := databases.Connect()

	tx := DB.Table(t.TableName()).Where("email = ?", t.Email).Updates(&User{
		EmailConfirmed: t.EmailConfirmed,
	})

	if tx.Error != nil {
		return errors.New("data not found")
	}

	return
}

func (t *User) EditPassword() (err error) {

	DB := databases.Connect()

	tx := DB.Debug().Table(t.TableName()).Where("email = ?", t.Email).Updates(&User{
		Password: t.Password,
	})

	if tx.Error != nil {
		return errors.New("data not found")
	}

	return
}

func (t *User) DeleteUser() (err error) {

	DB := databases.Connect()

	tx := DB.Table(t.TableName()).Delete(&t)

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	if tx.Error != nil {
		return errors.New("delete failed")
	}
	return
}
