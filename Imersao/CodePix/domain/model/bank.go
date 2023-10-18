package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
  govalidator.SetFieldsRequiredByDefault(true)
}

// NOTE:Struct can be serialized to JSON
type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name     string     `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	// NOTE: valid:"-" => indicates that the field can be null
	Accounts []*Account `gorm:"ForeignKey:BankID" valid:"-"`
}

// NOTE: attach this method to Bank struct
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	
	if err != nil {
		return err
	}
	
	return nil
}

// "Constructor": generate a struct
// NOTE: * => pointer (value by reference)
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	// nil = blank error (empty)
	// NOTE: & => return a reference
	return &bank, nil
}
