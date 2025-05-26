// Do not communicate by sharing memory; instead, share memory by communicating.

// concurrency safe
package bank 

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {deposits <- amount}
func Balance() {return <- balances}

func teller(){
	var balance int 
	for{
		select{
		case amount := <-deposits:
			balance +=amount
		case balances <- balance:
		}
	}
}

// this starts automatically on import
func init(){
	go teller()
}


