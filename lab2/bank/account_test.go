package bank

import "testing"

//Тест успешного создания счёта и пополнения
func TestDeposit(t *testing.T) {
	acc, err := NewAccount(100)
	if err != nil {
		t.Errorf("Не удалось создать счёт: %v", err)
	}

	err = acc.Deposit(50)
	if err != nil {
		t.Errorf("Неожиданая ошибка при выполнении: %f", err)
	}

	if acc.Balance() != 150 {
		t.Errorf("Ожидали баланс 150, но получили %f", acc.Balance())
	}

}

//тест успешного снятия средств
func TestWithdrawSuccess(t *testing.T) {
	acc, _ := NewAccount(100)

	err := acc.Withdraw(40)
	if err != nil {
		t.Errorf("Неожиданная ошибка при снятии: %v", err)
	}

	if acc.Balance() != 60 {
		t.Errorf("Ожидали баланс 60, но получили %f", acc.Balance())
	}
}

//Тест снять больше чем есть

func TestWithdrawMore(t *testing.T) {
	acc, _ := NewAccount(100)

	err := acc.Withdraw(150)
	if err == nil {
		t.Errorf("Ожидали ошибку о нехватке средств, но получили nil")
	}
}

//Тест передачи отрицательного
func TestWithdrawNegative(t *testing.T) {
	acc, _ := NewAccount(100)
	err := acc.Withdraw(-20)
	if err == nil {
		t.Errorf("Ожидали ошибку при снятии отрицательной суммы, но получили nil")
	}
}
