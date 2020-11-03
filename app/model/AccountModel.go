package model

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"implementasi-mvc/app/utils"
)

type AccountModel struct {
	DB *gorm.DB
	ID int				`gorm:"primary_key" json:"-"`
	IdAccount string	`json:"id_account,omitempty"`
	Name string			`json:"name"`
	Password string		`json:"password,omitempty"`
	AccountNumber int	`json:"account_number,omitempty"`
	Saldo int			`json:"saldo"`
}

func (account AccountModel) InsertNewAccount() (bool, error) {
	account.AccountNumber = utils.RangeIn(1000, 99999999)
	account.Saldo = 0
	account.IdAccount = fmt.Sprintf("id-%d", utils.RangeIn(10, 5000))

	result := account.DB.Create(&account)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (account AccountModel) GetAccountDetail(idAccount int) (bool, error, []Transaction, AccountModel) {
	var transaction []Transaction

	result := account.DB.Where("sender = ? OR recipient = ?", idAccount, idAccount).
		Find(&transaction)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return  false, errors.Errorf("Account not found"), []Transaction{}, AccountModel{}
		}

		return  false, result.Error, []Transaction{}, AccountModel{}
	}

	result = account.DB.Where(&AccountModel{
		AccountNumber: idAccount,
	}).Find(&account)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return  false, errors.Errorf("Account not found"), []Transaction{}, AccountModel{}
		}

		return  false, result.Error, []Transaction{}, AccountModel{}
	}

	return  true, nil, transaction, account
}


