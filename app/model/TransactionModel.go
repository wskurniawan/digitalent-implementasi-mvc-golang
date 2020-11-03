package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"implementasi-mvc/app/constant"
	"time"
)

type Transaction struct {
	ID int `gorm:"primary_key" json:"-"`
	TransactionType int `json:"transaction_type,omitempty"`
	TransactionDescription string `json:"transaction_description"`
	Sender int `json:"sender"`
	Amount int `json:"amount"`
	Recipient int `json:"recipient"`
	Timestamp int64 `json:"timestamp,omitempty"`
}

type TransactionModel struct {
	DB *gorm.DB
}

func (model TransactionModel) Transfer(trx Transaction) (bool, error) {
	err := model.DB.Transaction(func (tx *gorm.DB) error{
		var sender, recipient Account
		result := tx.Model(&Account{}).Where(&Account{
			AccountNumber: trx.Sender,
		}).First(&sender)
		if result.Error != nil {
			return  result.Error
		}

		// Check balance first
		if sender.Saldo < trx.Amount {
			return errors.Errorf("Insufficient saldo")
		}
		result = result.Update("saldo", sender.Saldo - trx.Amount)
		if result.Error != nil {
			return  result.Error
		}

		result = tx.Model(&Account{}).Where(Account{
			AccountNumber: trx.Recipient,
		}).First(&recipient).
			Update("saldo", recipient.Saldo + trx.Amount)
		if result.Error != nil {
			return  result.Error
		}

		trx.TransactionType = constant.TRANSFER
		trx.Timestamp = time.Now().Unix()
		result = tx.Create(&trx)
		if result.Error != nil {
			return  result.Error
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (model TransactionModel) Withdraw(trx Transaction) (bool, error) {
	err := model.DB.Transaction(func(tx *gorm.DB) error {
		var sender Account
		result := tx.Model(&Account{}).Where(&Account{
			AccountNumber: trx.Sender,
		}).First(&sender)
		if result.Error != nil {
			return  result.Error
		}

		// Check balance first
		if sender.Saldo < trx.Amount {
			return errors.Errorf("Insufficient saldo")
		}

		result = result.Update("saldo", sender.Saldo - trx.Amount)
		if result.Error != nil {
			return  result.Error
		}

		trx.TransactionType = constant.WITHDRAW
		trx.Timestamp = time.Now().Unix()
		result = tx.Create(&trx)
		if result.Error != nil {
			return  result.Error
		}

		return  nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (model TransactionModel) Deposit(trx Transaction) (bool, error) {
	err := model.DB.Transaction(func (tx *gorm.DB) error {
		var sender Account
		result := tx.Model(&Account{}).Where(&Account{
			AccountNumber: trx.Sender,
		}).First(&sender).
			Update("saldo", sender.Saldo + trx.Amount)

		if result.Error != nil {
			return result.Error
		}

		trx.TransactionType = constant.DEPOSIT
		trx.Timestamp = time.Now().Unix()
		result = tx.Create(&trx)
		if result.Error != nil {
			return result.Error
		}

		return  nil
	})

	if err != nil {
		return false, err
	}

	return  true, nil
}