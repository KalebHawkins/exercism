package account

import (
	"sync"
)

// Define the Account type here.
type Account struct {
	m                sync.Mutex
	AvailableBalance int64
	IsClosed         bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	acc := Account{}
	acc.AvailableBalance += amount

	return &acc
}

func (a *Account) Balance() (int64, bool) {
	a.m.Lock()
	defer a.m.Unlock()

	if a.IsClosed {
		return 0, false
	}

	return a.AvailableBalance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.m.Lock()
	defer a.m.Unlock()

	if !a.IsClosed && a.AvailableBalance+amount >= 0 {
		a.AvailableBalance += amount
		return a.AvailableBalance, true
	}

	return a.AvailableBalance, false
}

func (a *Account) Close() (int64, bool) {
	a.m.Lock()
	defer a.m.Unlock()

	if a.IsClosed {
		return 0, !a.IsClosed
	}

	a.IsClosed = true
	return a.AvailableBalance, a.IsClosed
}
