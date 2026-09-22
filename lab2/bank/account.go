package bank

import (
	"errors"
)

// банковский счёт с балансом
type BankAccount struct {
	balance float64
}

// создаём новый счёт
func NewAccount(initialBalance float64) (*BankAccount, error) {
	if initialBalance < 0 {
		return nil, errors.New("Начальный баланс не может быть отрицательным")
	}
	return &BankAccount{balance: initialBalance}, nil
}

// текущий баланс
func (a *BankAccount) Balance() float64 {
	return a.balance
}

// пополняем баланс
func (a *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Сумма поплнения должна быть больше нуля")
	}
	a.balance += amount
	return nil
}

func (a *BankAccount) Withdraw(amoant float64) error {
	if amoant <= 0 {
		return errors.New("Сумма снятия должна быть больше нуля")
	}
	if amoant > a.balance {
		return errors.New("Недостаточно средств на счёте")
	}
	a.balance -= amoant
	return nil
}
