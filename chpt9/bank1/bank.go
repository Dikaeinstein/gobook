// Package bank provides a concurrency-safe bank with one account.
package bank

type withdrawAmount struct {
	amount float64
	c      chan bool
}

var deposits = make(chan float64)           // send amount to deposit
var balances = make(chan float64)           // receive balance
var withdrawals = make(chan withdrawAmount) // send amount to withdraw

// Deposit adds amount to your current account balance
func Deposit(amount float64) { deposits <- amount }

// Balance returns the current account balance
func Balance() float64 { return <-balances }

// Withdraw amount from account
func Withdraw(amount float64) bool {
	status := make(chan bool)
	w := withdrawAmount{amount, status}
	withdrawals <- w
	return <-status
}

func teller() {
	var balance float64
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawals:
			result := balance - w.amount
			if result > 0 {
				balance = result
				w.c <- true // Operation succeeded
			} else {
				w.c <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
