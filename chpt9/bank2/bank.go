package bank

var (
	sema    = make(chan struct{}, 1) // Binary semaphore
	balance float64
)

// Deposit adds amount to your current account balance
func Deposit(amount float64) {
	sema <- struct{}{} // acquire token
	balance += amount
	<-sema // release token
}

// Balance returns the current account balance
func Balance() float64 {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

// Withdraw amount from account
func Withdraw(amount float64) bool {
	var r bool
	b := balance - amount
	sema <- struct{}{}
	if b > 0 {
		balance = b
		r = true
	} else {
		r = false
	}
	<-sema
	return r
}
