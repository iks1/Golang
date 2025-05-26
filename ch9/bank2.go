var (
	sema = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int){
	sema <- struct{}{} // acquire token
	balance = balance + amount 
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{}
	b := balance 
	<-sema 
	return b
}