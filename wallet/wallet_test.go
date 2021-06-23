package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		AssertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(1))
		AssertBalance(t, wallet, Bitcoin(19))
		AssertNoError(t, err)
	})

	t.Run("Withdraw insufficent funds", func(t *testing.T) {
		initialBalance := Bitcoin(2)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(82))
		AssertBalance(t, wallet, initialBalance)
		AssertError(t, err, ErrInsufficientFunds)
	})
}
