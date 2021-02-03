import "time"

package model

import (
"github.com/asaskevich/govalidator"
uuid "github.com/satori/go.uuid"
"time"
)

type User struct {
	Base     `valid:"required"`
	Code     string     `json:"code" valid:"notnull"`
	Name     string     `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(code string, name string) (*User, error) {
	user := User{
		Code: code,
		Name: name,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
