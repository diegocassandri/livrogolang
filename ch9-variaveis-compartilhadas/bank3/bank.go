// bank provides a concurrency_safe sing;e-account bank
package bank

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	defer mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
