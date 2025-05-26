var (
	mu sync.RWmutex
	balance int
)

// this allows multiple go routines to read, which is more practical 
func Balance() int {
	mu.RLock() // readers lock 
	defer mu.RUnlock()
	return balance 
}

// Deposit function is unchanged