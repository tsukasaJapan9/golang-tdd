package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertMsg := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		t.Helper()
		if got == nil {
			t.Error("wanted error")
		}

		if got.Error() != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Error("wanted no error")
		}
	}

	t.Run("Deposite", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(Bitcoin(10))
		assertMsg(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance:Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))
		assertMsg(t, wallet, Bitcoin(0))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient balance", func(t *testing.T) {
		wallet := Wallet{balance:Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertMsg(t, wallet, Bitcoin(20))
		assertError(t, err, "mount is short")
	})
}