package bank

import "sync"

var (
	balance float64
	mu      sync.RWMutex // Guards balance
)

// Deposit adds amount to your current account balance
func Deposit(amount float64) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

// This function requires that the lock be held.
func deposit(amount float64) {
	balance += amount
}

// Balance returns the current account balance
func Balance() float64 {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

// Withdraw amount from account
func Withdraw(amount float64) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount) // Deduct amount
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
