import "sync"

var (
	mu sync.Mutex
	balance int 
)

func Deposit(amount int){
	mu.Lock()
	nalance = balance + amount 
	mu.Unlock()
}

// not so good implementation
func Balance() int {
	mu.Lock() 
	b := balance 
	mu.Unlock()
	return b
}
// better implementation
func Balance1() int {
	mu.Lock()
	defer mu.Unlock()
	return balance 
}

// consider withdraw function 

//Note: not atomic
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

//incorrect implementation 
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

//correct implementation
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if Balance() < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {balance += amount}
