package repository

import (
	"fmt"
	"github.com/alexandrecpedro/fullcycle/Imersao/CodePix/go/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (t *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	// NOTE: brings AccountFrom and Bank related with the transaction 
	// Passes the value to variable 'transaction', where the condition is satisfied
	// Avoiding SQL Injection
	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
