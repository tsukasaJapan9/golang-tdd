package wallet

import (
	"fmt"
	"github.com/pkg/errors"
)

var errorInsufficient = errors.New("mount is short")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposite(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errorInsufficient
	}

	w.balance -= amount
	return nil
}


