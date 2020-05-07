package account

import "sync"

// Account struct.
type Account struct {
	balance int
	opened  bool
	mtx     sync.Mutex // Protects balance from read/writes
}

// Open receives an balance (int) and returns a new Account
func Open(balance int) *Account {
	if balance < 0 {
		return nil
	}

	return &Account{balance: balance, opened: true}
}

// Balance takes no parameters and returns an int and a bool.
func (a Account) Balance() (int, bool) {
	return a.balance, a.opened
}

// Close takes no parameters and returns an int and a bool.
func (a *Account) Close() (balance int, opened bool) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	{ // sync unlocked
		opened = a.opened
		if a.opened {
			balance = a.balance
			a.opened = false
			a.balance = 0
		}
	}
	return
}

// Deposit accepts an balance (balance) and returns an int and a bool.
func (a *Account) Deposit(balance int) (int, bool) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	{ // sync locked
		if a.opened {
			if 0 > (balance + a.balance) {
				return 0, false
			}

			a.balance += balance
		}
	}

	return a.balance, a.opened
}
