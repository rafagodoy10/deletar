package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Cliente struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func NewCliente(kind string, account *Account, key string) (*Cliente, error) {

	cliente := Cliente{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	Cliente.ID = uuid.NewV4().String()
	Cliente.CreatedAt = time.Now()

	err := Cliente.isValid()
	if err != nil {
		return nil, err
	}

	return &Cliente, nil
}